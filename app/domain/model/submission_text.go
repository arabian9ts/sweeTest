package model

import "time"

type SubmissionText struct {
	ID           int64
	SubmissionID int64
	FileName     int64
	Contents     string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SubmissionTexts []*SubmissionText
