package usecase

import "github.com/kikils/golang-todo/domain/model"

type UserRepository interface {
	Store(model.User) (int, error)
	Update(user model.User) (id int, err error)
	Delete(userID int) (err error)
	FindById(int) (model.User, error)
	FindAll() (model.Users, error)
}
