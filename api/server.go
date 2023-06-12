package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/hyeonjungko/bankofasia/db/sqlc"
)

// Server serves HTTP requests
type Server struct {
	store  db.Store
	router *gin.Engine
}

// NewServer creates new server and sets up routing
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)
	router.DELETE("/accounts/:id", server.deleteAccount)
	router.PATCH("/accounts", server.updateAccount)

	server.router = router
	return server
}

// Start runs HTTP server on given address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
