package dto

import "time"

//input
type CreateStudentLectureInputForm struct {
	StudentID int64 `validate:"required,gt=0" json:"-"`
	LectureID int64 `validate:"required,gt=0" json:"-"`
}

type CreateAssistantLectureInputForm struct {
	StudentID int64 `validate:"required,gt=0" json:"-"`
	LectureID int64 `validate:"required,gt=0" json:"-"`
}

type CreateTeacherLectureInputForm struct {
	TeacherID int64 `validate:"required,gt=0" json:"-"`
	LectureID int64 `validate:"required,gt=0" json:"-"`
}

//output
type CreateStudentLectureOutputForm struct {
	ID        int64     `json:"id"`
	StudentID int64     `validate:"required,gt=0" json:"student_id"`
	LectureID int64     `validate:"required,gt=0" json:"lecture_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateAssistantLectureOutputForm struct {
	ID        int64     `json:"id"`
	StudentID int64     `validate:"required,gt=0" json:"student_id"`
	LectureID int64     `validate:"required,gt=0" json:"lecture_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateTeacherLectureOutputForm struct {
	ID        int64     `json:"id"`
	TeacherID int64     `validate:"required,gt=0" json:"teacher_id"`
	LectureID int64     `validate:"required,gt=0" json:"lecture_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type DeleteStudentLectureOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}

type DeleteAssistantLectureOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}

type DeleteTeacherLectureOutputForm struct {
	AffectedRowsCount int64 `json:"count"`
}
