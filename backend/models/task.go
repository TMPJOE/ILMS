package models

import "time"

type Status int

const (
	created Status = iota
	inProgress
	pending
	done
)

type Task struct {
	Id     int       `json:"id"`
	Status Status    `json:"status"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}

type TaskInput struct {
	Name string `json:"name"`
}
