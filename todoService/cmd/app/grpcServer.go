package app

import (
	"net"

	"github.com/mohammaderm/todoMicroService/todoService/config"
	delivery_grpc "github.com/mohammaderm/todoMicroService/todoService/internal/delivery/grpc"
	"github.com/mohammaderm/todoMicroService/todoService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/todoService/proto"

	"google.golang.org/grpc"
)

func Server(cfg config.Grpc, usecase usecase.TodoService) error {
	lis, err := net.Listen("tcp", ":"+cfg.Port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()

	todoGrpc := delivery_grpc.NewTodoDelivery(usecase)
	categoryGrpc := delivery_grpc.NewCategoryDelivery(usecase)

	proto.RegisterTodoServiceServer(s, todoGrpc)
	proto.RegisterCategoryServiceServer(s, categoryGrpc)

	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
