package dto

import "time"

// Input
type CreateHelpInputForm struct {
	LectureID int64  `validate:"required,gt=0"     json:"-"`
	StudentID int64  `validate:"required,gt=0"     json:"-"`
	Contents  string `validate:"required,lte=1000" json:"contents"`
}

type UpdateHelpInputForm struct {
	ID        int64  `validate:"required,gt=0"     json:"-"`
	LectureID int64  `validate:"required,gt=0,"    json:"-"`
	StudentID int64  `validate:"required,gt=0"     json:"-"`
	Contents  string `validate:"required,lte-1000" json:"contents"`
}

// Output
type GetHelpOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	StudentID int64     `json:"student_id"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetHelpsOutputForm []*GetHelpOutputForm

type CreateHelpOutputForm struct {
	LastChangedHelpID int64 `json:"last_changed_help_id"`
}

type UpdateHelpOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteHelpOutputForm struct {
	Deleted bool `json:"deleted"`
}
