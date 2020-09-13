module github.com/kikils/golang-todo

go 1.14

require (
	github.com/lib/pq v1.8.0
)

replace (
	github.com/kikils/golang-todo/model => ./domain/model
	github.com/kikils/golang-todo/infrastructure => ./infrastructure
	github.com/kikils/golang-todo/controllers => ./interfaces/controllers
	github.com/kikils/golang-todo/database => ./interfaces/database
	github.com/kikils/golang-todo/usecase => ./usecase
)