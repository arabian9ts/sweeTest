package model

import "time"

type Class struct {
	ID        int64
	LectureID int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Classes []*Class
