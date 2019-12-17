package model

import "time"

type Task struct {
	ID        int64
	ClassID   int64
	Title     string
	Desc      string
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tasks []*Task
