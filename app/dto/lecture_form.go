package dto

import "time"

// Input
type CreateLectureInputForm struct {
	Name string `validate:"required,lte=50" json:"name"`
}

type UpdateLectureInputForm struct {
	ID   int64  `validate:"required,gt=0"`
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

type CreateLectureOutputForm struct {
	LastChangedLectureId int64 `json:"changed_lecture_id"`
}

type UpdateLectureOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteLectureOutputForm struct {
	Deleted bool `json:"deleted"`
}
