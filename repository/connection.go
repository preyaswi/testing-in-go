package repository

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connectDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=preya dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
