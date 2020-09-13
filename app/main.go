package main

import (
	"log"
	"net/http"

	"github.com/kikils/golang-todo/infrastructure"
)

func main() {
	mux := infrastructure.SetUpRouting()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
