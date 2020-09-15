package database

import (
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
		return
	}
	for row.Next() {
		if err := row.Scan(&id); err != nil {
			return -1, err
		}
	}
	return
}

func (repo *UserRepository) Update(user model.User) (id int, err error) {
	row, err := repo.Sqlhandler.Query(
		"UPDATE users SET FirstName=$1, LastName=$2 WHERE id=$3 RETURNING id;", user.FirstName, user.LastName, user.ID,
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

func (repo *UserRepository) Delete(userID int) (err error) {
	_, err = repo.Sqlhandler.Query("DELETE FROM users WHERE id=$1", userID)
	if err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindById(userID int) (user model.User, err error) {
	row, err := repo.Sqlhandler.Query("SELECT id, FirstName, LastName FROM users WHERE id = $1;", userID)
	if err != nil {
		return
	}
	row.Next()
	if err = row.Scan(&(user.ID), &(user.FirstName), &(user.LastName)); err != nil {
		return
	}
	return
}

func (repo *UserRepository) FindAll() (users model.Users, err error) {
	rows, err := repo.Sqlhandler.Query("SELECT id, FirstName, LastName FROM users;")
	if err != nil {
		return
	}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&(user.ID), &(user.FirstName), &(user.LastName)); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}
