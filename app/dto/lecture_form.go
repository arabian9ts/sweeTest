package dto

import "time"

type WriteLectureInputForm struct {
	Name string
}

type ReadLectureOutputForm struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ReadLecturesOutputForm []*ReadLectureOutputForm

type WriteLectureOutputForm struct {
	LastChangedLectureId int64
}
