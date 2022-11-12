package grpc

import (
	"github.com/mohammaderm/todoMicroService/todoService/internal/models"
	"github.com/mohammaderm/todoMicroService/todoService/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapTodos(todos *[]models.Todo) []*proto.Todo {
	protoTodos := []*proto.Todo{}
	for _, todo := range *todos {
		protoTodos = append(protoTodos, &proto.Todo{
			Id:          todo.Id,
			AccountId:   todo.AccountId,
			CategoryId:  todo.CategoryId,
			Title:       todo.Title,
			Description: todo.Description,
			Status:      todo.Status,
			CreatedAt:   timestamppb.New(todo.CreatedAt),
		})
	}
	return protoTodos
}

func mapCategorys(categorys *[]models.Category) []*proto.Category {
	protoCategorys := []*proto.Category{}
	for _, category := range *categorys {
		protoCategorys = append(protoCategorys, &proto.Category{
			Id:        category.Id,
			Title:     category.Title,
			AccountId: category.AccountId,
			CreatedAt: timestamppb.New(category.CreatedAt),
		})
	}
	return protoCategorys
}
