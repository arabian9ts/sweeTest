package dto

import "time"

// Input
type CreateSubmissionInputForm struct {
	TaskID    int64 `json:"-" validate:"required,gt=0"`
	StudentID int64 `json:"-" validate:"required,gt=0"`
}

type CreateSubmissionTextInputForm struct {
	FileName string `json:"file_name" validate:"required,lte=50"`
	Contents string `json:"contents"`
}

type UpdateSubmissionTextInputForm struct {
	ID           int64  `json:"-" validate:"required,gt=0"`
	SubmissionID int64  `json:"-" validate:"required,gt=0"`
	FileName     string `json:"file_name" validate:"required,lte=50"`
	Contents     string `json:"contents"`
}

// Output
type GetSubmissionOutputForm struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	StudentID int64     `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetSubmissionsOutputForm []*GetSubmissionOutputForm

type GetSubmissionTextOutputForm struct {
	ID           int64     `json:"id"`
	SubmissionID int64     `json:"submission_id"`
	FileName     string    `json:"file_name"`
	Contents     string    `json:"contents"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type GetSubmissionTextsOutputForm []*GetSubmissionTextOutputForm

type CreateSubmissionOutputForm struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	StudentID int64     `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateSubmissionTextOutputForm struct {
	ID           int64     `json:"id"`
	SubmissionID int64     `json:"submission_id"`
	FileName     string    `json:"file_name"`
	Contents     string    `json:"contents"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UpdateSubmissionOutputForm struct {
	ID        int64     `json:"id"`
	TaskID    int64     `json:"task_id"`
	StudentID int64     `json:"student_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateSubmissionTextOutputForm struct {
	ID           int64     `json:"id"`
	SubmissionID int64     `json:"submission_id"`
	FileName     string    `json:"file_name"`
	Contents     string    `json:"contents"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type DeleteSubmissionTextOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
