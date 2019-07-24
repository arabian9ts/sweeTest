package dto

import "time"

// Input
type CreateLectureInputForm struct {
	Name string
}

type UpdateLectureInputForm struct {
	Name string
}

// Output
type GetLectureByIdOutputForm struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetLecturesOutputForm []*GetLectureByIdOutputForm

type CreateLectureOutputForm struct {
	LastChangedLectureId int64
}

type UpdateLectureOutputForm struct {
	Updated bool
}

type DeleteLectureOutputForm struct {
	Deleted bool
}
