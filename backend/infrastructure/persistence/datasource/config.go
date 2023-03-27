package datasource

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connection() (*sqlx.DB,error) {
	username :=os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")
	url := fmt.Sprintf("%s:%s@%s:3306/general?charset=utf8", username, password, host)
	if username == "" {
		url = "docker:docker@tcp(localhost:3306)/general?charset=utf8"
	}
	db, err := sqlx.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	return db, nil
}