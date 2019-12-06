package model

import "time"

type Help struct {
	ID        int64
	LectureID int64
	StudentID int64
	Title     string
	Contents  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Helps []*Help
