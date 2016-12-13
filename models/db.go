package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func getDBConnection() (*sql.DB, error) {
	db, err := sql.Open("postgres", "dbname=go_blog user=francescorenzi host=localhost port=5432 sslmode=disable")
	return db, err
}
