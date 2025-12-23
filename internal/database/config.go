package database

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
)

func NewConnection() {
	dbuser := os.Getenv("DB_USERNAME")
	dbpassword := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://%s:%s@localhost/%s", dbuser, dbpassword, dbname))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println()

	log.Println("DB CONNECTED")
}
