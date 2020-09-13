package infrastructure

import (
	"database/sql"
	"fmt"

	"github.com/kikils/golang-todo/interfaces/database"
	_ "github.com/lib/pq" // postres driver
)

type Sqlhandler struct {
	DB *sql.DB
}

func NewSqlhandler() *Sqlhandler {
	connStr := "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil
	}

	err = db.Ping()
	if err != nil {
		fmt.Println("here")
		return nil
	}

	return &Sqlhandler{db}
}

func (handler *Sqlhandler) Execute(statement string, args ...interface{}) (database.Result, error) {
	res := SqlResult{}
	result, err := handler.DB.Exec(statement, args...)
	if err != nil {
		return res, err
	}
	res.Result = result
	return res, nil
}

func (handler *Sqlhandler) Query(statement string, args ...interface{}) (database.Row, error) {
	rows, err := handler.DB.Query(statement, args...)
	if err != nil {
		return new(SqlRow), err
	}
	row := new(SqlRow)
	row.Rows = rows
	return row, nil
}

type SqlResult struct {
	Result sql.Result
}

func (r SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) error {
	return r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) Close() error {
	return r.Rows.Close()
}
