package models

import "time"

type Task struct {
	Id     int       `json:"id" db:"id"`
	Status int       `json:"status" db:"status"`
	Name   string    `json:"name" db:"name"`
	Desc   string    `json:"desc" db:"description"`
	Date   time.Time `json:"date" db:"created_at"`
}

type TaskInput struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type TaskOutput struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	Date   string `json:"date"`
}

type TaskUpdate struct {
	Id     int    `json:"id"`
	Status int    `json:"status"`
	Name   string `json:"name"`
	Desc   string `json:"desc"`
}
