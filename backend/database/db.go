package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func OpenDb() (*pgxpool.Pool, error) {
	dbpool, err := pgxpool.New(context.Background(), os.Getenv("DB_URL"))
	if err != nil {
		return dbpool, nil
	}

	return &pgxpool.Pool{}, err
}
