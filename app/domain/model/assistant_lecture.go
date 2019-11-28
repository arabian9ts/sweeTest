package model

import "time"

type AssistantLecture struct {
	ID        int64
	StudentID int64
	LectureID int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type AssistantLectures []*AssistantLecture
