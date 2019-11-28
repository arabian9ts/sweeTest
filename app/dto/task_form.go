package dto

import "time"

// Input
type CreateTaskInputForm struct {
	LectureID int64     `validate:"required,gt=0" json:"-"`
	Title     string    `validate:"lte=100" json:"title"`
	Desc      string    `json:"desc"`
	Deadline  time.Time `json:"deadline"`
}

type UpdateTaskInputForm struct {
	ID        int64     `validate:"required,gt=0" json:"-"`
	LectureID int64     `validate:"required,gt=0" json:"-"`
	Title     string    `validate:"required,lte=100" json:"title"`
	Desc      string    `json:"desc"`
	Deadline  time.Time `json:"deadline"`
}

// Output
type GetTaskOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetTasksByLectureIdOutputForm []*GetTaskOutputForm

type CreateTaskOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UpdateTaskOutputForm struct {
	ID        int64     `json:"id"`
	LectureID int64     `json:"lecture_id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteTaskOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
