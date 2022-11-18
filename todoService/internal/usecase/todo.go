package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mohammaderm/todoMicroService/todoService/internal/dto"
	"github.com/mohammaderm/todoMicroService/todoService/internal/models"
	"github.com/mohammaderm/todoMicroService/todoService/internal/repository"
	"github.com/mohammaderm/todoMicroService/todoService/pkg/validator"
)

type (
	Service struct {
		repo  repository.TodoRepository
		cache todoCacheInterface
	}

	TodoService interface {
		// todo
		Create(ctx context.Context, req dto.CreateTodoReq) error
		GetAll(ctx context.Context, req dto.GetAllTodoReq) (dto.GetAllTodoRes, error)
		Delete(ctx context.Context, req dto.DeleteTodoReq) error
		Update(ctx context.Context, req dto.UpdateTodoReq) error

		// category
		CreateCat(ctx context.Context, req dto.CreateCatReq) error
		GetAllCat(ctx context.Context, req dto.GetAllCatReq) (dto.GetAllCatRes, error)
		DeleteCat(ctx context.Context, req dto.DeleteCatReq) error
	}
)

func New(repo repository.TodoRepository, cache todoCacheInterface) TodoService {
	return &Service{
		repo:  repo,
		cache: cache,
	}
}

// -----------------
// category
// -----------------
func (s *Service) DeleteCat(ctx context.Context, req dto.DeleteCatReq) error {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("input is not valid: %w", err)
	}
	err = s.repo.DeleteCat(ctx, req.Id, req.AccountId)
	if err != nil {
		return err
	}
	err = s.cache.deleteAll(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetAllCat(ctx context.Context, req dto.GetAllCatReq) (dto.GetAllCatRes, error) {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return dto.GetAllCatRes{}, fmt.Errorf("input is not valid: %w", err)
	}
	result, err := s.cache.getAllCat(ctx, fmt.Sprintf("userCategory-%d", req.AccountId))
	if err != nil {
		result, err := s.repo.GetAllCat(ctx, req.AccountId)
		if err != nil {
			return dto.GetAllCatRes{}, err
		}
		err = s.cache.setAllCat(ctx, fmt.Sprintf("userCategory-%d", req.AccountId), result, time.Duration(30)*time.Minute)
		if err != nil {
			log.Printf("Warnig: failed to cache valuse: %s \n", err.Error())
		}
		return dto.GetAllCatRes{
			Categorys: result,
		}, nil
	}
	return dto.GetAllCatRes{
		Categorys: result,
	}, nil
}

func (s *Service) CreateCat(ctx context.Context, req dto.CreateCatReq) error {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("input is not valid: %w", err)
	}
	var category models.Category
	category.AccountId = req.AccountId
	category.Title = req.Title
	err = s.repo.CreateCat(ctx, &category)
	if err != nil {
		return err
	}
	s.cache.deleteAll(ctx)
	return nil
}

// -----------------
// todo
// -----------------
func (s *Service) Create(ctx context.Context, req dto.CreateTodoReq) error {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("input is not valid: %w", err)
	}
	var todo models.Todo
	todo.AccountId = req.AccountId
	todo.CategoryId = req.CategoryId
	todo.Title = req.Title
	todo.Description = req.Description
	err = s.repo.Create(ctx, &todo)
	if err != nil {
		return err
	}
	s.cache.deleteAll(ctx)
	return nil
}

func (s *Service) GetAll(ctx context.Context, req dto.GetAllTodoReq) (dto.GetAllTodoRes, error) {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return dto.GetAllTodoRes{}, fmt.Errorf("input is not valid: %w", err)
	}
	result, err := s.cache.getAll(ctx, fmt.Sprintf("userTodo-%d-ofset-%d", req.AccountId, req.Offset))
	if err != nil {
		result, err = s.repo.GetAll(ctx, req.AccountId, req.Offset)
		if err != nil {
			return dto.GetAllTodoRes{}, err
		}
		err = s.cache.setAll(ctx, fmt.Sprintf("userTodo-%d-ofset-%d", req.AccountId, req.Offset), result, time.Duration(30)*time.Minute)
		if err != nil {
			log.Printf("Warnig: failed to cache valuse: %s \n", err.Error())
		}
		return dto.GetAllTodoRes{
			Todos: result,
		}, nil
	}
	println("cache hit")
	return dto.GetAllTodoRes{
		Todos: result,
	}, nil
}

func (s *Service) Delete(ctx context.Context, req dto.DeleteTodoReq) error {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("input is not valid: %w", err)
	}
	err = s.repo.Delete(ctx, req.Id, req.AccountId)
	if err != nil {
		return err
	}
	err = s.cache.deleteAll(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) Update(ctx context.Context, req dto.UpdateTodoReq) error {
	err := validator.TodoRequest(ctx, req)
	if err != nil {
		return fmt.Errorf("input is not valid: %w", err)
	}
	var todo models.Todo
	todo.AccountId = req.AccountId
	todo.CategoryId = req.CategoryId
	todo.Description = req.Description
	todo.DueDate = req.DueDate
	todo.Id = req.Id
	todo.Priority = req.Priority
	todo.Status = req.Status
	todo.Title = req.Title

	err = s.repo.Update(ctx, &todo)
	if err != nil {
		return err
	}
	err = s.cache.deleteAll(ctx)
	if err != nil {
		return err
	}
	return nil
}
