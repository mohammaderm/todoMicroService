package grpc

import (
	"context"

	"github.com/mohammaderm/todoMicroService/todoService/internal/dto"
	"github.com/mohammaderm/todoMicroService/todoService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/todoService/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CategoryServer struct {
	todoUsecase usecase.TodoService
	proto.UnimplementedCategoryServiceServer
}

func NewCategoryDelivery(todoUsecase usecase.TodoService) proto.CategoryServiceServer {
	return &CategoryServer{
		todoUsecase: todoUsecase,
	}
}

// ------------------------------
// -----category-------
// ------------------------------

func (c *CategoryServer) DeleteCat(ctx context.Context, req *proto.DeleteCatRequest) (*proto.DeleteCatRespons, error) {
	request := dto.DeleteCatReq{
		Id:        uint(req.Id),
		AccountId: uint(req.AccountId),
	}
	err := c.todoUsecase.DeleteCat(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.DeleteCatRespons{
		Error:   false,
		Message: "category deleted succesfully",
	}, nil
}

func (c *CategoryServer) CreateCat(ctx context.Context, req *proto.CreateCatRequest) (*proto.CreateCatRespons, error) {
	request := dto.CreateCatReq{
		Title:     req.Title,
		AccountId: req.AccountId,
	}
	result, err := c.todoUsecase.CreateCat(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CreateCatRespons{
		Error:   false,
		Message: "category sucesfully created",
		Category: &proto.Category{
			Id:        result.Category.Id,
			Title:     result.Category.Title,
			AccountId: request.AccountId,
			CreatedAt: timestamppb.New(result.Category.CreatedAt),
		},
	}, nil
}

func (c *CategoryServer) GetAllCat(ctx context.Context, req *proto.GetAllCatRequest) (*proto.GetAllCatRespons, error) {
	request := dto.GetAllCatReq{
		AccountId: uint(req.AccountId),
	}
	result, err := c.todoUsecase.GetAllCat(ctx, request)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.GetAllCatRespons{
		Categorys: mapCategorys(result.Categorys),
		Error:     false,
		Message:   "succesfully returned",
	}, err
}
