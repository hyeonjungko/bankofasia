package gapi

import (
	"fmt"

	db "github.com/hyeonjungko/bankofasia/db/sqlc"
	"github.com/hyeonjungko/bankofasia/pb"
	"github.com/hyeonjungko/bankofasia/token"
	"github.com/hyeonjungko/bankofasia/util"
)

// Server serves gRPC requests
type Server struct {
	pb.UnimplementedBankOfAsiaServer
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
}

// NewServer creates new gRPC server
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

	return server, nil
}
