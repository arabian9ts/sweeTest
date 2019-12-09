package dto

import "time"

// Input
type CreateClassInputForm struct {
	LectureID int64  `validate:"required,gt=0"   json:"-"`
	Name      string `validate:"required,lte=50" json:"name"`
}

type UpdateClassInputForm struct {
	ID        int64  `validate:"required,gt=0"   json:"-"`
	LectureID int64  `validate:"required,gt=0"   json:"-"`
	Name      string `validate:"required,lte=50" json:"name"`
}

// Output
type GetClassByIdOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetClassesOutputForm []*GetClassByIdOutputForm

type CreateClassOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateClassOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteClassOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
