package usecase

import "github.com/kikils/golang-todo/app/domain/model"

type UserInteractor struct {
	UserRepository UserRepository
}

func (interactor *UserInteractor) Add(u model.User) (id int, err error) {
	id, err = interactor.UserRepository.Store(u)
	return
}

func (interactor *UserInteractor) Users() (user model.Users, err error) {
	user, err = interactor.UserRepository.FindAll()
	return
}

func (interactor *UserInteractor) UserById(identifier int) (user model.User, err error) {
	user, err = interactor.UserRepository.FindById(identifier)
	return
}
