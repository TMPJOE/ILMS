package repositories

import (
	"database/sql"

	"dev.acevedo/backend/models"
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

func (t *TaskRepo) SelectAll() ([]*models.TaskOutput, error) {
	query := "SELECT * FROM task"

	rows, err := t.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*models.TaskOutput

	for rows.Next() {
		t := &models.TaskOutput{}

		err = rows.Scan(
			&t.Id,
			&t.Status,
			&t.Name,
			&t.Desc,
			&t.Date,
		)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, t)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return tasks, nil

}
