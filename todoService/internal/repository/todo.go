package repository

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mohammaderm/todoMicroService/todoService/internal/models"
)

var (
	ErrorNotFound         = errors.New("can not founf any todo with this id")
	ErrUniquenessViolated = errors.New("requested operation violates uniqueness of a field")
)

type repository struct {
	db *sqlx.DB
}

type TodoRepository interface {
	// todo
	Create(ctx context.Context, todo *models.Todo) (*models.Todo, error)
	Delete(ctx context.Context, id, accountid uint) error
	GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Todo, error)
	Update(ctx context.Context, todo *models.Todo) error

	// category
	CreateCat(ctx context.Context, category *models.Category) (*models.Category, error)
	DeleteCat(ctx context.Context, id, accountid uint) error
	GetAllCat(ctx context.Context, accountid uint) (*[]models.Category, error)
}

func New(db *sqlx.DB) TodoRepository {
	return &repository{
		db: db,
	}
}

// category
func (r *repository) GetAllCat(ctx context.Context, accountid uint) (*[]models.Category, error) {
	var result []models.Category
	err := r.db.SelectContext(ctx, &result, getAllCategory, accountid)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) DeleteCat(ctx context.Context, id, accountid uint) error {
	result, err := r.db.ExecContext(ctx, deleteCategory, id, accountid)
	if err != nil {
		return err
	}
	rowEfected, _ := result.RowsAffected()
	if rowEfected == 0 {
		return ErrorNotFound
	}
	return nil
}

func (r *repository) CreateCat(ctx context.Context, category *models.Category) (*models.Category, error) {
	var id int
	err := r.db.QueryRowContext(ctx, createCategory, category.Title, category.AccountId).Scan(&id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return nil, fmt.Errorf("%w: an category with the given title already exists", ErrUniquenessViolated)
		}
	}
	var catResult models.Category
	err = r.db.GetContext(ctx, &catResult, getCategoryById, strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	return &catResult, nil
}

// todo

func (r *repository) Update(ctx context.Context, todo *models.Todo) error {
	result, err := r.db.ExecContext(ctx, updateTodo, todo.Title, todo.CategoryId, todo.Description, todo.Status, todo.DueDate, todo.Priority, todo.Id, todo.AccountId)
	if err != nil {
		return err
	}
	rowEfected, _ := result.RowsAffected()
	if rowEfected == 0 {
		return ErrorNotFound
	}
	return nil
}

func (r *repository) GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Todo, error) {
	var result []models.Todo
	err := r.db.SelectContext(ctx, &result, getAllTodo, accountid, limit, offset)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r *repository) Delete(ctx context.Context, id, accountid uint) error {
	result, err := r.db.ExecContext(ctx, deleteTodo, id, accountid)
	if err != nil {
		return err
	}
	rowEfected, _ := result.RowsAffected()
	if rowEfected == 0 {
		return ErrorNotFound
	}
	return nil
}

func (r *repository) Create(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	var id int
	err := r.db.QueryRowContext(ctx, createTodo, todo.Title, todo.Description, todo.CategoryId, todo.AccountId).Scan(&id)
	if err != nil {
		return nil, err
	}
	var todoResult models.Todo
	err = r.db.GetContext(ctx, &todoResult, getTodoById, strconv.Itoa(id))
	if err != nil {
		return nil, err
	}
	return &todoResult, nil
}
