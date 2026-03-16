package services

import (
	"log/slog"

	"dev.acevedo/backend/models"
	"dev.acevedo/backend/repositories"
)

type TaskService struct {
	l *slog.Logger
	r *repositories.TaskRepo
}

func NewTaskService(l *slog.Logger, r *repositories.TaskRepo) *TaskService {
	return &TaskService{
		l: l,
		r: r,
	}
}

func (s *TaskService) AddTask(input models.TaskInput) (models.TaskOutput, error) {
	result, err := s.r.Create(input.Name, input.Desc)
	if err != nil {
		s.l.Error("Something went wrong", "desc", err.Error())
		return models.TaskOutput{}, err
	}

	s.l.Info(
		"Task Added",
		slog.String("Name", input.Name),
	)

	return *result, nil
}

func (s *TaskService) GetTasks(id int) (*models.TaskResponse, error) {

	var lastId int

	if id < 0 {
		id = 0
	}
	tasks, err := s.r.Select(id)
	if err != nil {
		s.l.Error("Something broke", "err", err)
		return nil, err
	}

	//allocate memory for task output slice
	tasksOut := make([]models.TaskOutput, 0, len(tasks))

	//iterate through slice of pointers to assign its value to the slice of task output
	for _, task := range tasks {
		tasksOut = append(tasksOut, *task)
		lastId = task.Id
	}

	response := &models.TaskResponse{
		Tasks:  tasksOut,
		LastId: lastId,
	}

	s.l.Info("Showing all tasks")

	return response, nil
}

func (s *TaskService) UpdateTask(update models.TaskUpdate) (models.TaskSwapBack, error) {
	task, err := s.r.Update(update)
	if err != nil {
		s.l.Error("Update error caused by", "err", err)
		return models.TaskSwapBack{}, err
	}

	s.l.Info("Task updated")

	taskOut := *task

	return taskOut, err
}

func (s *TaskService) DeleteTask(id int) error {
	result, err := s.r.Delete(id)
	rows, _ := result.RowsAffected()
	if err != nil {
		s.l.Error("Delete error caused by", "err", err, slog.Int64("Rows affected", rows))
		return err
	}

	s.l.Info("Task deleted successfully", slog.Int64("Rows affected", rows))
	return nil
}
