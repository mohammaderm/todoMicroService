package app

import (
	"net"

	"github.com/mohammaderm/authService/config"
	delivery_grpc "github.com/mohammaderm/authService/internal/delivery/grpc"
	"github.com/mohammaderm/authService/internal/usecase"
	"github.com/mohammaderm/authService/pkg/logger"
	"github.com/mohammaderm/authService/proto"
	"google.golang.org/grpc"
)

func Server(cfg config.Grpc, usecase usecase.AuthUscase, logger logger.Logger) error {
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return err
	}
	server := grpc.NewServer()

	authGrpc := delivery_grpc.New(logger, usecase)
	proto.RegisterAuthServiceServer(server, authGrpc)

	if err := server.Serve(lis); err != nil {
		return err
	}
	return nil
}
