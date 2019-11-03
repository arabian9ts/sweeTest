package model

import "time"

type TeacherLecture struct {
	ID        int64
	TeacherID int64
	LectureID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type TeacherLectures []*TeacherLecture
