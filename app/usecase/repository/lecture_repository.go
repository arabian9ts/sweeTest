package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type LectureRepository interface {
	GetParticipatedLecturesOfStudent(studentID int64, limit int, offset int) (model.Lectures, error)
	GetParticipatedLecturesOfAssistant(studentID int64, limit int, offset int) (model.Lectures, error)
	GetParticipatedLecturesOfTeacher(teacherID int64, limit int, offset int) (model.Lectures, error)

	GetLectures(limit int, offset int) (model.Lectures, error)
	GetLectureById(id int64) (*model.Lecture, error)
	CreateLecture(lecture *model.Lecture) (*model.Lecture, error)
	UpdateLecture(lecture *model.Lecture) (*model.Lecture, error)
	DeleteLecture(id int64) (int64, error)
}
