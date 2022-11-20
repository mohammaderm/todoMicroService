package types

import "time"

type RegisterReq struct {
	UserName string `json:"username" example:"example5040"`
	Email    string `json:"email" example:"example@gmai.com"`
	Password string `json:"password" example:"111222333444"`
}

type LoginReq struct {
	Email    string `json:"email" example:"example@gmai.com"`
	Password string `json:"password" example:"111222333444"`
}

type LoginRes struct {
	Account      Account `json:"account"`
	AccessToken  string  `json:"accessToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTY5Nzc2MjEsImZuYW1lIjoiU2hhbiIsImxuYW1lIjoiVml2IiwidXNlciI6ImFzZEB0ZXN0LmNvbSJ9.tdhUL-KpDmzSNtV9z6XhUgoTKcVabuOPS3fHAySjSXQ"`
	RefreshToken string  `json:"refreshToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE1OTY5Nzc2MjEsImZuYW1lIjoiU2hhbiIsImxuYW1lIjoiVml2IiwidXNlciI6ImFzZEB0ZXN0LmNvbSJ9.tdhUL-KpDmzSNtV9z6XhUgoTKcVabuOPS3fHAySjSXQ"`
}

type AccountInfo struct {
	Id uint64
}

type Account struct {
	Id        uint64    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
