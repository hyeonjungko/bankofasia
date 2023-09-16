package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	db "github.com/hyeonjungko/bankofasia/db/sqlc"
	"github.com/hyeonjungko/bankofasia/token"
	"github.com/hyeonjungko/bankofasia/util"
)

// Server serves HTTP requests
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates new HTTP server and sets up routing
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// users
	router.POST("/users", server.createUser)

	// user login
	router.POST("/users/login", server.loginUser)
	router.POST("tokens/renew_access", server.renewAccessToken)

	// Create and Add Auth Middleware for rest of routes
	authRoutes := router.Group("/").Use(authMiddleWare(server.tokenMaker))

	// accounts
	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.DELETE("/accounts/:id", server.deleteAccount)
	authRoutes.PATCH("/accounts", server.updateAccount)

	// transfers
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs HTTP server on given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
