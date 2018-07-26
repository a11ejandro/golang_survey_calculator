package main

import (
	"os"
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

var db *sql.DB
var err error

func connectToDb() {
	user := os.Getenv("POSTGRES_USERNAME")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DATABASE")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user, password, dbname, host, port)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
	}

	if err = db.Ping(); err != nil {
			 fmt.Println(err)
	 }
}
