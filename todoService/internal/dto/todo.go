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
	UpdateTodoStatusReq struct {
		Id        uint `json:"id" validate:"required"`
		AccountId uint `json:"accountid" validate:"required"`
	}
	UpdateTodoPriorityReq struct {
		Id        uint `json:"id" validate:"required"`
		AccountId uint `json:"accountid" validate:"required"`
		Priority  int  `json:"priority" db:"priority" validate:"required"`
	}
	UpdateTodoDueDateReq struct {
		Id        uint      `json:"id" validate:"required"`
		AccountId uint      `json:"accountid" validate:"required"`
		DueDate   time.Time `json:"due_date" db:"due_date"`
	}
	GetAllTodoRes struct {
		Todos *[]models.Todo `json:"todos"`
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
)
