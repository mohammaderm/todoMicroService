package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"github.com/mohammaderm/authService/internal/models"
	"github.com/mohammaderm/authService/pkg/logger"
)

var (
	ErrNotFound           = errors.New("not found")
	ErrUniquenessViolated = errors.New("requested operation violates uniqueness of a field")
)

type Repository struct {
	logger logger.Logger
	db     *sqlx.DB
}

type UserRepository interface {
	Create(ctx context.Context, user *models.User) error
	FindById(ctx context.Context, id uint) (*models.User, error)
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

func New(logger logger.Logger, db *sqlx.DB) UserRepository {
	return &Repository{
		logger: logger,
		db:     db,
	}
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, &result, GetUserbyEmailQuery, email)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("an user with email = %s was not fount: %w", email, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user account: %w", err)
	}
	return &result, nil
}

func (r *Repository) FindById(ctx context.Context, id uint) (*models.User, error) {
	var result models.User
	err := r.db.GetContext(ctx, &result, GetUserbyIdQuery, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("an user with id = %d was not fount: %w", id, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch user account: %w", err)
	}
	return &result, nil
}

func (r *Repository) Create(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(ctx, CreateUserQuery, user.Username, user.Email, user.Password)
	if err != nil {
		r.logger.Error("failed to insert user", map[string]interface{}{
			"error": err.Error(),
		})
		if err, ok := err.(*pq.Error); ok && err.Code.Name() == "unique_violation" {
			return fmt.Errorf("%w: an account with the given email already exists", ErrUniquenessViolated)
		}
	}
	return nil
}
