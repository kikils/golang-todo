package infrastructure

import (
	"log"
	"github.com/kikils/golang-todo/interfaces"
	"database/sql"

	_ "github.com/lib/pq" // postres driver
)

type Sqlhandler struct {
	DB *sql.DB
}

func ConnectPostgres() (*Sqlhandler, error) {
	connStr := "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Sqlhandler{db}, nil
}

func (handler *Sqlhandler) Execute(statement string, args ...interface{}) error {
	_, err := handler.DB.Exec(statement, args...)
	return err
}

func (handler *Sqlhandler) Query(statement string) interfaces.Row {
	//fmt.Println(statement)
	rows, err := handler.DB.Query(statement)
	if err != nil {
		log.Fatal(err)
		return new(SqlRow)
	}
	row := new(SqlRow)
	row.Rows = rows
	return row
}

type SqlRow struct {
	Rows *sql.Rows
}

func (r SqlRow) Scan(dest ...interface{}) {
	r.Rows.Scan(dest...)
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}