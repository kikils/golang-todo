package usecase

import (
	"context"
	"github.com/kikils/golang-todo/domain"
)

const keyInteractor = "Interactor"

type Interactor interface {
	UserInteractor
	TodoInteractor
}

type UserInteractor interface {
    AddUser(user *domain.User) (*domain.User, error)
	DeleteUser(id int) (error)
}

func AddUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return getInteractor(ctx).AddUser(user)
}

func DeleteUser(ctx context.Context, id int) error {
	return getInteractor(ctx).DeleteUser(id)
}

type TodoInteractor interface {
    AddTodo(todo *domain.Todo) (*domain.Todo, error)
	DeleteTodo(id int) error
	GetAllTodo() ([]domain.Todo, error)
}

func AddTodo(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	return getInteractor(ctx).AddTodo(todo)
}

func DeleteTodo(ctx context.Context, id int) error {
	return getInteractor(ctx).DeleteTodo(id)
}

func GetAllTodo(ctx context.Context) ([]domain.Todo, error) {
	return getInteractor(ctx).GetAllTodo()
}

func SetUserInteractor(ctx context.Context, interactor Interactor) context.Context {
	return context.WithValue(ctx, keyInteractor, interactor)
}

func getInteractor(ctx context.Context) Interactor {
	return ctx.Value(keyInteractor).(Interactor)
}

