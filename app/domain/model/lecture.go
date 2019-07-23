package model

import "time"

type Lecture struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Lectures []*Lecture
