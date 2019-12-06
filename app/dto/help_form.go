package dto

import "time"

// Input
type CreateHelpInputForm struct {
	LectureID int64  `validate:"required,gt=0"     json:"-"`
	StudentID int64  `validate:"required,gt=0"     json:"-"`
	Title     string `validate:"required,lte=1000" json:"title"`
	Contents  string `validate:"required,lte=1000" json:"contents"`
}

type UpdateHelpInputForm struct {
	ID        int64  `validate:"required,gt=0"     json:"-"`
	LectureID int64  `validate:"required,gt=0"     json:"-"`
	StudentID int64  `validate:"required,gt=0"     json:"-"`
	Title     string `validate:"required,lte=1000" json:"title"`
	Contents  string `validate:"required,lte=1000" json:"contents"`
}

// Output
type GetHelpOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	StudentID int64     `json:"student_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetHelpsOutputForm []*GetHelpOutputForm

type CreateHelpOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	StudentID int64     `json:"student_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateHelpOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	StudentID int64     `json:"student_id"`
	Title     string    `json:"title"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteHelpOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
