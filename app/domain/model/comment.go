package model

import "time"

type Comment struct {
	ID        int64
	HelpID    int64
	UserType  UserType
	UserID    int64
	Contents  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Comments []*Comment
