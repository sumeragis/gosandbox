package datasource

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connection() (*sqlx.DB,error) {
    url := os.Getenv("MYSQL_DATASOURCE_URL")
	if url == "" {
		url = "docker:docker@tcp(localhost:3306)/general?charset=utf8"
	}

	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}