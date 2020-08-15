package domain

import (
	"time"
)

type Todo struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Note    string    `json:"note"`
	DueDate time.Time `json:"due_date"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
