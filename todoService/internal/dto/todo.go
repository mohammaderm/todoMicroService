package dto

import (
	"time"

	"github.com/mohammaderm/todoMicroService/todoService/internal/models"
)

type (
	// todo
	CreateTodoReq struct {
		AccountId   uint64 `json:"accountid" db:"accountid" validate:"required"`
		CategoryId  uint64 `json:"categoryid" db:"categoryid" validate:"required"`
		Title       string `json:"title" db:"title" validate:"required"`
		Description string `json:"description" db:"description"`
	}
	GetAllTodoReq struct {
		AccountId uint `json:"accountid" validate:"required"`
		Offset    int  `json:"offset"`
	}
	DeleteTodoReq struct {
		Id        uint `json:"id" validate:"required"`
		AccountId uint `json:"accountid" validate:"required"`
	}

	UpdateTodoReq struct {
		Id          uint64    `json:"id" validate:"number"`
		AccountId   uint64    `json:"accountid" validate:"number"`
		CategoryId  uint64    `json:"categoryid" validate:"number"`
		Title       string    `json:"title"`
		Description string    `json:"description"`
		Status      bool      `json:"status" validate:"boolean"`
		DueDate     time.Time `json:"due_date" validate:"datetime"`
		Priority    int       `json:"priority" validate:"number"`
	}

	GetAllTodoRes struct {
		Todos *[]models.Todo `json:"todos"`
	}
	CreateTodoRes struct {
		Todo *models.Todo `json:"todo"`
	}

	GetAllByCtegoryReq struct {
		AccountId  uint64 `json:"accountid" db:"accountid" validate:"required"`
		CategoryId uint64 `json:"categoryid" db:"categoryid" validate:"required"`
	}

	// category
	CreateCatReq struct {
		Title     string `json:"title" db:"title" validate:"required"`
		AccountId uint64 `json:"accountid" db:"accountid" validate:"required"`
	}
	GetAllCatReq struct {
		AccountId uint `json:"accountid" validate:"required"`
	}
	DeleteCatReq struct {
		Id        uint `json:"id" validate:"required"`
		AccountId uint `json:"accountid" validate:"required"`
	}
	GetAllCatRes struct {
		Categorys *[]models.Category `json:"categorys"`
	}
	CreateCatRes struct {
		Category *models.Category `json:"categorys"`
	}
)
