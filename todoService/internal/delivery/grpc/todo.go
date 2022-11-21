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

type TodoServer struct {
	todoUsecase usecase.TodoService
	proto.UnimplementedTodoServiceServer
}

func NewTodoDelivery(todoUsecase usecase.TodoService) proto.TodoServiceServer {
	return &TodoServer{
		todoUsecase: todoUsecase,
	}
}

// ------------------------------
// -------todo-------
// ------------------------------

func (t *TodoServer) Update(ctx context.Context, req *proto.UpdateRequest) (*proto.UpdateRespons, error) {
	todo := dto.UpdateTodoReq{
		Id:          req.Id,
		AccountId:   req.AccountId,
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Description: req.Description,
		Status:      req.Status,
		DueDate:     req.DueDate.AsTime(),
		Priority:    int(req.Priority),
	}
	err := t.todoUsecase.Update(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateRespons{
		Error:   false,
		Message: "succesfully updated",
	}, nil
}

func (t *TodoServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteRespons, error) {
	todo := dto.DeleteTodoReq{
		Id:        uint(req.Id),
		AccountId: uint(req.AccountId),
	}
	err := t.todoUsecase.Delete(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.DeleteRespons{
		Error:   false,
		Message: "succesfully deleted",
	}, nil
}

func (t *TodoServer) GetAll(ctx context.Context, req *proto.GetAllRequest) (*proto.GetAllRespons, error) {
	todo := dto.GetAllTodoReq{
		AccountId: uint(req.AccountId),
		Offset:    int(req.Offset),
	}
	todos, err := t.todoUsecase.GetAll(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.GetAllRespons{
		Todos:   mapTodos(todos.Todos),
		Error:   false,
		Message: "succesfully returned",
	}, nil
}

func (t *TodoServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateRespons, error) {
	todo := dto.CreateTodoReq{
		AccountId:   req.AccountId,
		CategoryId:  req.CategoryId,
		Title:       req.Title,
		Description: req.Description,
	}
	result, err := t.todoUsecase.Create(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CreateRespons{
		Error:   false,
		Message: "todo succesfully created",
		Todo: &proto.Todo{
			Id:          result.Todo.Id,
			AccountId:   result.Todo.AccountId,
			CategoryId:  result.Todo.CategoryId,
			Title:       result.Todo.Title,
			Description: result.Todo.Description,
			Status:      result.Todo.Status,
			CreatedAt:   timestamppb.New(result.Todo.CreatedAt),
			Priority:    uint64(result.Todo.Priority),
			DueDate:     timestamppb.New(result.Todo.DueDate),
		},
	}, nil

}

func (t *TodoServer) GetAllByCategoryId(ctx context.Context, req *proto.GetAllByCategoryIdRequest) (*proto.GetAllByCategoryIdRespons, error) {
	todo := dto.GetAllByCtegoryReq{
		AccountId:  req.AccountId,
		CategoryId: req.CategoryId,
	}
	todos, err := t.todoUsecase.GetAllByCategory(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.GetAllByCategoryIdRespons{
		Todos:   mapTodos(todos.Todos),
		Error:   false,
		Message: "succesfully returned",
	}, nil
}
