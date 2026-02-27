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

func (s *TaskService) AddTask(input models.TaskInput) error {
	result, err := s.r.Create(input.Name, input.Desc)
	if err != nil {
		s.l.Error("Something went wrong", "desc", err.Error())
		return err
	}

	id, _ := result.LastInsertId()

	s.l.Info(
		"Task Added",
		slog.String("Name", input.Name),
		slog.Int64("id", id),
	)

	return nil
}
