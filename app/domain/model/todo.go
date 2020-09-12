package model

import "time"

type Todo struct {
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Note    string    `json:"note"`
	DueDate time.Time `json:"due_date"`
	UserID  int       `json:"user_id"`
}

type Todos []Todo
