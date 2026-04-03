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

func (t *TaskRepo) Create(taskName, taskDesc string) (*models.TaskOutput, error) {
	query := "INSERT INTO task (status, name, description, created_at) VALUES (0,?,?, CURRENT_TIMESTAMP)"

	result, err := t.db.Exec(query, taskName, taskDesc)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	query = "SELECT * FROM task WHERE id = ?"
	row := t.db.QueryRow(query, id)
	task := &models.TaskOutput{}

	if err = row.Scan(
		&task.Id,
		&task.Status,
		&task.Name,
		&task.Desc,
		&task.Date,
	); err != nil {
		return nil, err
	}

	return task, nil
}

func (t *TaskRepo) Select(id int) ([]*models.TaskOutput, error) {
	query := "SELECT * FROM task WHERE (id) > (?) ORDER BY id, created_at DESC LIMIT 11;"

	rows, err := t.db.Query(query, id)
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

func (t *TaskRepo) Update(task models.TaskUpdate) (*models.TaskSwapBack, error) {
	query := `
	UPDATE task 
	SET name = ?, 
		status = ?, 
		description = ? 
	WHERE id = ?;`

	_, err := t.db.Exec(query, task.Name, task.Status, task.Desc, task.Id)
	if err != nil {
		return nil, err
	}

	query = "SELECT name, description FROM task WHERE id = ?"
	row := t.db.QueryRow(query, task.Id)
	taskBack := &models.TaskSwapBack{}
	if err = row.Scan(&taskBack.Name, &taskBack.Desc); err != nil {
		return nil, err
	}

	return taskBack, nil
}

func (t *TaskRepo) Delete(id int) (sql.Result, error) {
	query := "DELETE from task WHERE id = ?"

	result, err := t.db.Exec(query, id)
	if err != nil {
		return result, err
	}

	return result, nil
}
