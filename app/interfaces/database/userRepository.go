package database

import (
	"log"

	"github.com/kikils/golang-todo/domain/model"
)

type UserRepository struct {
	Sqlhandler
}

func (repo *UserRepository) Store(u model.User) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"INSERT INTO users (FirstName, LastName) VALUES ($1,$2) RETURNING id;", u.FirstName, u.LastName,
	)

	if err != nil {
		log.Println(err)
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *UserRepository) FindById(identifier int) (user model.User, err error) {
	row, err := repo.Sqlhandler.Query("SELECT id, FirstName, LastName FROM users WHERE id = $1;", identifier)
	defer row.Close()
	if err != nil {
		log.Println(err)
		return
	}
	var id int
	var firstName string
	var lastName string
	row.Next()
	if err = row.Scan(&id, &firstName, &lastName); err != nil {
		return
	}
	user.ID = id
	user.FirstName = firstName
	user.LastName = lastName
	return
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT id, FirstName, LastName FROM users")
	defer rows.Close()
	if err != nil {
		return
	}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		if err := rows.Scan(&id, &firstName, &lastName); err != nil {
			continue
		}
		user := model.User{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		users = append(users, user)
	}
	return
}
