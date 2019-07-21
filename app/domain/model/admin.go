package model

import "time"

type Admin struct {
	ID        int64
	FirstName string
	LastName  string
	Email     string
	Digest    string
	CreatedAt time.Time
	UpdatedAt time.Time
}
