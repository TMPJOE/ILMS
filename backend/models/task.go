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
	id     int
	status Status
	name   string
	date   time.Time
}
