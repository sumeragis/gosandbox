package datasource

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() (*sql.DB,error) {
	db, err := sql.Open("mysql", "docker:docker@tcp(localhost:3306)/general?charset=utf8")
	if err != nil {
		return nil, err
	}
	return db, nil
}