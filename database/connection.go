package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)


func ConnectDB() {
	var dsn string
	if dsn = os.Getenv("MYSQL_CONNECTION_STRING"); dsn == "" {
		panic("MySQL connection string hasn't been provided.")
	}

	db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}