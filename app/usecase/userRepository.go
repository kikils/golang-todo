package usecase

import "github.com/kikils/golang-todo/domain/model"

type UserRepository interface {
	Store(model.User) (int, error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
}
