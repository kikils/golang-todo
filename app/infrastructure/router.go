package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kikils/golang-todo/interfaces/controllers"
)

func SetUpRouting() *http.ServeMux {
	mux := http.NewServeMux()

	sqlhandler := NewSqlhandler()
	err := CreateTable(sqlhandler)
	if err != nil {
		log.Println(err.Error())
	}
	userController := controllers.NewUserController(sqlhandler)
	todoController := controllers.NewTodoController(sqlhandler)

	mux.HandleFunc("/user/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Create(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Show(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/user/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Update(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/user/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Delete(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/todo/create", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Create(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Show(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/update", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Update(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/delete", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Delete(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/search", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Search(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	return mux
}

func ResponseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")

	body := map[string]string{
		"error": message,
	}
	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
