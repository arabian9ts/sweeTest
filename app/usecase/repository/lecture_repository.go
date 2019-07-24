package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type LectureRepository interface {
	GetLectures(limit int, offset int) (model.Lectures, error)
	GetLectureById(id int64) (*model.Lecture, error)
	CreateLecture(lecture *model.Lecture) (int64, error)
	UpdateLecture(lecture *model.Lecture) (int64, error)
	DeleteLecture(id int64) (int64, error)
}
