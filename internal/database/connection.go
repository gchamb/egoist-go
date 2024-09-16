package database

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func ConnectDB() (*sqlx.DB) {
	var dsn string
	if dsn = os.Getenv("MYSQL_CONNECTION_STRING"); dsn == "" {
		panic("MySQL connection string hasn't been provided.")
	}

	db, err := sqlx.Connect("mysql", dsn)
    if err != nil {
        panic(err)
    }
    fmt.Println("Connected!")
	return db
}