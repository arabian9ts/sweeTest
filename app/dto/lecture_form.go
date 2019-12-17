package dto

import (
	"time"
)

// Input
type CreateLectureInputForm struct {
	Name string `validate:"required,lte=50" json:"name"`
}

type UpdateLectureInputForm struct {
	ID   int64  `validate:"required,gt=0" json:"-"`
	Name string `validate:"required,lte=50" json:"name"`
}

// Output
type GetLectureByIdOutputForm struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetLecturesOutputForm []*GetLectureByIdOutputForm

type GetTotalLecturesOutputForm struct {
	Total    int64                 `json:"total"`
	Lectures GetLecturesOutputForm `json:"lectures"`
}

type CreateLectureOutputForm struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateLectureOutputForm struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteLectureOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
