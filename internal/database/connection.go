package database

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


func ConnectDB() (*sqlx.DB) {
	var dsn string
	if dsn = os.Getenv("MYSQL_CONNECTION_STRING"); dsn == "" {
		panic("MySQL connection string hasn't been provided.")
	}

	db, err := sqlx.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
	
    fmt.Println("Connected!")
	return db
}