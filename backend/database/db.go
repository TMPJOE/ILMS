package database

import (
	"database/sql"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "backend/database/tasks.db")
	if err != nil {
		return db, nil
	}

	return &sql.DB{}, err
}
