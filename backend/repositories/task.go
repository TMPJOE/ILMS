package repositories

import (
	"database/sql"
)

type TaskRepo struct {
	db *sql.DB
}

func NewTaskRepo(dbConnection *sql.DB) *TaskRepo {
	return &TaskRepo{
		db: dbConnection,
	}
}

func (t *TaskRepo) Create(taskName, taskDesc string) (sql.Result, error) {
	query := "INSERT INTO task (status, name, description, created_at) VALUES (0,?,?, CURRENT_TIMESTAMP)"

	result, err := t.db.Exec(query, taskName, taskDesc)
	if err != nil {
		return result, err
	}

	return result, nil
}
