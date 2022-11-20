package dto

import "github.com/mohammaderm/todoMicroService/authService/internal/models"

type RegisterReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=16"`
}

type RegisterRes struct {
	Message string `json:"message"`
	Error   bool   `json:"error"`
}

type LoginReq struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type LoginRes struct {
	User         *models.User `json:"user"`
	AccessToken  string       `json:"accesstoken"`
	RefreshToken string       `json:"refreshtoken"`
}
