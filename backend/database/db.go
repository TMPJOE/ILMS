package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "backend/database/tasks.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}
