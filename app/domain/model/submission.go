package model

import "time"

type Submission struct {
	ID             int64
	TaskID         int64
	StudentID      int64
	SubmissionText SubmissionText
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Submissions []*Submission
