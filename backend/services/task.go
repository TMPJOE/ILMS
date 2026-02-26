package services

import (
	"log/slog"

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

func (s *TaskService) AddTask(taskName string) error {
	_, err := s.r.Create(taskName)
	if err != nil {
		s.l.Error("Something went wrong", "desc", err.Error())
		return err
	}

	return nil
}
