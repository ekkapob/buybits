package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	DB_USER = "rugbit"
	DB_PWD  = "rugbit_123"
	DB_NAME = "rugbitdb"
)

func getDb() *sql.DB {
	connection := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PWD, DB_NAME)
	db, err := sql.Open("postgres", connection)
	panicIfError(err)
	return db
}
