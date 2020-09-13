package infrastructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/kikils/golang-todo/interfaces/controllers"
)

func SetUpRouting() *http.ServeMux {
	// router := gin.Default()
	mux := http.NewServeMux()

	sqlhandler := NewSqlhandler()
	userController := controllers.NewUserController(sqlhandler)
	todoController := controllers.NewTodoController(sqlhandler)

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userController.Create(w, r)
		case http.MethodGet:
			userController.Index(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/user/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			userController.Show(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})

	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			todoController.Create(w, r)
		case http.MethodGet:
			todoController.Index(w, r)
		default:
			ResponseError(w, http.StatusNotFound, "")
		}
	})
	mux.HandleFunc("/todo/get", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			todoController.Show(w, r)
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
