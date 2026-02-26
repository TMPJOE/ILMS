package repositories

import (
	"database/sql"
)

type TaskRepo struct {
	db *sql.DB
}

func (t *TaskRepo) dbConstruct(dbConnection *sql.DB) {

	t.db = dbConnection
}

func NewTaskRepo(dbConnection *sql.DB) *TaskRepo {
	return &TaskRepo{
		db: dbConnection,
	}
}

func (t *TaskRepo) Create(taskName string) (sql.Result, error) {
	query := "INSERT INTO task (status, taskName, date) VALUE (0,?, CURRENT_TIMESTAMP)"

	result, err := t.db.Exec(query, taskName)
	if err != nil {
		return result, err
	}

	return result, nil
}
