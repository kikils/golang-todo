package interfaces

import (
	"github.com/kikils/golang-todo/usecase"
	"github.com/kikils/golang-todo/domain"
	"context"
)

type DbHandler interface {
	Execute(statement string, args ...interface{}) error
	Query(statement string) Row
}

type Row interface {
	Scan(dest ...interface{})
	Next() bool
}

type Repository interface {
	Execute(statement string, args ...interface{}) error
	Query(statement string) Row
}

func AddTodo(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	return usecase.AddTodo(ctx, todo)
}

func DeleteTodo(ctx context.Context, id int) error {
	return usecase.DeleteTodo(ctx, id)
}

func GetAllTodo(ctx context.Context) ([]domain.Todo, error) {
	return usecase.GetAllTodo(ctx)
}