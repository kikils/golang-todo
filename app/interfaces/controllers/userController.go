package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/kikils/golang-todo/domain/model"
	"github.com/kikils/golang-todo/interfaces/database"
	"github.com/kikils/golang-todo/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.Sqlhandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				Sqlhandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var user model.User
	if err := json.Unmarshal(b, &user); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	id, err := controller.Interactor.Add(user)
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseOk(w, id)
}

func (controller *UserController) Index(w http.ResponseWriter, r *http.Request) {
	users, err := controller.Interactor.Users()
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseOk(w, users)
}

func (controller *UserController) Show(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		ID int `json:"id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := controller.Interactor.UserById(req.ID)
	if err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}
	responseOk(w, user)
}
