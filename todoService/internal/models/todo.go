package models

import "time"

type Todo struct {
	Id          uint64    `json:"id" db:"id"`
	AccountId   uint64    `json:"accountid" db:"accountid"`
	CategoryId  uint64    `json:"categoryid" db:"categoryid"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	Status      bool      `json:"status" db:"status"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	DueDate     time.Time `json:"due_date" db:"due_date"`
	Priority    int       `json:"priority" db:"priority"`
}

type Category struct {
	Id        uint64    `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	AccountId uint64    `json:"accountid" db:"accountid"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}
