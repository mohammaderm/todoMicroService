package repository

import (
	"context"
	"errors"
	"fmt"
	"time"

	my "github.com/go-mysql/errors"
	"github.com/jmoiron/sqlx"
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
	Create(ctx context.Context, todo *models.Todo) error
	Delete(ctx context.Context, id, accountid uint) error
	GetAll(ctx context.Context, accountid uint, offset int) (*[]models.Todo, error)
	UpdateStatus(ctx context.Context, id, accountid uint) error
	UpdatePriority(ctx context.Context, id, accountid uint, priority int) error
	UpdateDueDate(ctx context.Context, id, accountid uint, due_date time.Time) error

	// category
	CreateCat(ctx context.Context, category *models.Category) error
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

func (r *repository) CreateCat(ctx context.Context, category *models.Category) error {
	_, err := r.db.ExecContext(ctx, createCategory, category.Title, category.AccountId)
	if err != nil {
		if ok, err := my.Error(err); ok && err == my.ErrDupeKey {
			return fmt.Errorf("%w: an category with the given title already exists", ErrUniquenessViolated)
		}
	}
	return nil
}

// todo
func (r *repository) UpdateStatus(ctx context.Context, id, accountid uint) error {
	_, err := r.db.ExecContext(ctx, updateStatus, id, accountid)
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) UpdatePriority(ctx context.Context, id, accountid uint, priority int) error {
	_, err := r.db.ExecContext(ctx, updatePriority, priority, id, accountid)
	if err != nil {
		return err
	}
	return nil
}
func (r *repository) UpdateDueDate(ctx context.Context, id, accountid uint, due_date time.Time) error {
	_, err := r.db.ExecContext(ctx, updateDueDate, due_date, id, accountid)
	if err != nil {
		return err
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

func (r *repository) Create(ctx context.Context, todo *models.Todo) error {
	_, err := r.db.ExecContext(ctx, createTodo, todo.Title, todo.Description, todo.CategoryId, todo.AccountId)
	if err != nil {
		return err
	}
	return nil
}
