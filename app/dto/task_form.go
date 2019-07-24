package dto

import "time"

// Input
type CreateTaskInputForm struct {
	LectureID int64
	Title     string
	Desc      string
	Deadline  time.Time
}

type UpdateTaskInputForm struct {
	ID        int64
	LectureID int64
	Title     string
	Desc      string
	Deadline  time.Time
}

// Output
type GetTaskOutputForm struct {
	ID        int64
	LectureID int64
	Title     string
	Desc      string
	Deadline  time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type GetTasksByLectureIdOutputForm []*GetTaskOutputForm

type CreateTaskOutputForm struct {
	LastChangedTaskId int64
}

type UpdateTaskOutputForm struct {
	Updated bool
}

type DeleteTaskOutputForm struct {
	Deleted bool
}
