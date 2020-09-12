package usecase

import "github.com/kikils/golang-todo/app/domain/model"

type UserRepository interface {
	Store(model.User) (int, error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
}
