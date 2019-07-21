package model

import "time"

type Ta struct {
	ID        int64
	StudentNo string
	FirstName string
	LastName  string
	Email     string
	Digest    string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}
