package grpc

import (
	"context"

	"github.com/mohammaderm/todoMicroService/todoService/internal/dto"
	"github.com/mohammaderm/todoMicroService/todoService/internal/usecase"
	"github.com/mohammaderm/todoMicroService/todoService/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (t *TodoServer) UpdateDueDate(ctx context.Context, req *proto.UpdateDueDateRequest) (*proto.UpdateRespons, error) {
	todo := dto.UpdateTodoDueDateReq{
		Id:        uint(req.Id),
		AccountId: uint(req.AccountId),
		DueDate:   req.DueDate.AsTime(),
	}
	err := t.todoUsecase.UpdateDueDate(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateRespons{
		Error:   false,
		Message: "todo succesfully done",
	}, nil
}

func (t *TodoServer) UpdatePriority(ctx context.Context, req *proto.UpdatePriorityRequest) (*proto.UpdateRespons, error) {
	todo := dto.UpdateTodoPriorityReq{
		Id:        uint(req.Id),
		AccountId: uint(req.AccountId),
		Priority:  int(req.Priority),
	}
	err := t.todoUsecase.UpdatePriority(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateRespons{
		Error:   false,
		Message: "todo succesfully done",
	}, nil
}

func (t *TodoServer) UpdateStatus(ctx context.Context, req *proto.UpdateStatusRequest) (*proto.UpdateRespons, error) {
	todo := dto.UpdateTodoStatusReq{
		Id:        uint(req.Id),
		AccountId: uint(req.AccountId),
	}
	err := t.todoUsecase.UpdateStatus(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.UpdateRespons{
		Error:   false,
		Message: "todo succesfully done",
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
	err := t.todoUsecase.Create(ctx, todo)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &proto.CreateRespons{
		Error:   false,
		Message: "todo succesfully created",
	}, nil

}
