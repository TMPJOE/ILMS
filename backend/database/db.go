package database

import (
	"database/sql"
	"embed"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func OpenDb(dbFiles embed.FS, logger *slog.Logger, dbPath string) (*sql.DB, error) {
	logger.Info("Opening database", "path", dbPath)

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		logger.Error("Failed to open database", "error", err.Error())
		return nil, err
	}

	// Check if tables exist, if not run migration
	var tableCount int
	err = db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='task'").Scan(&tableCount)
	if err != nil {
		logger.Error("Failed to check database schema", "error", err.Error())
		return nil, err
	}

	if tableCount == 0 {
		logger.Info("No tables found, running migration")

		migrationSQL, err := dbFiles.ReadFile("sql/dbMigration.sql")
		if err != nil {
			return nil, fmt.Errorf("failed to read migration SQL: %w", err)
		}

		_, err = db.Exec(string(migrationSQL))
		if err != nil {
			logger.Error("Migration failed", "error", err.Error())
			return nil, fmt.Errorf("failed to run migration: %w", err)
		}

		logger.Info("Database migration completed")
	}

	logger.Info("Database connection established")
	return db, nil
}
