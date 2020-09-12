package database

import (
	"github.com/kikils/golang-todo/domain/model"
)

type TodoRepository struct {
	Sqlhandler
}

func (repo *TodoRepository) Store(t model.Todo) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"INSERT INTO todos (title, note, duedate, userid) VALUES ($1,$2,$3,$4) RETURNING id;", t.Title, t.Note, t.DueDate, t.UserID,
	)

	if err != nil {
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *TodoRepository) FindById(id int) (todo model.Todo, err error) {
	row, err := repo.Sqlhandler.Query("SELECT id, title, note, duedate FROM todos WHERE id = $1;", id)
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DueDate)); err != nil {
		return
	}
	return
}

func (repo *TodoRepository) FindByUserId(userID int) (todoList model.Todos, err error) {
	query := `SELECT DISTINCT todos.id, title, note, duedate FROM todos INNER JOIN users ON todos.userid=users.id AND users.id=$1;`
	rows, err := repo.Sqlhandler.Query(query, userID)
	if err != nil {
		return
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DueDate)); err != nil {
			continue
		}
		todoList = append(todoList, todo)
	}
	return
}

func (repo *TodoRepository) FindAll() (todoList model.Todos, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT id, title, note, duedate FROM todos;")
	if err != nil {
		return
	}
	for rows.Next() {
		var todo model.Todo
		if err := rows.Scan(&(todo.ID), &(todo.Title), &(todo.Note), &(todo.DueDate)); err != nil {
			continue
		}
		todoList = append(todoList, todo)
	}
	return
}
