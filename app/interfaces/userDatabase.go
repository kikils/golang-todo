package interfaces

import (
	"github.com/kikils/golang-todo/domain"
	"context"
	"github.com/kikils/golang-todo/usecase"
)

func AddUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	return usecase.AddUser(ctx, user)
}

func DeleteUser(ctx context.Context, id int) error {
	return usecase.DeleteUser(ctx, id)
}