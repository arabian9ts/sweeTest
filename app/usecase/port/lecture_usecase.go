package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type LectureUseCase interface {
	GetLectureUseCase
	CreateLectureUseCase
	UpdateLectureUseCase
	DeleteLectureUseCase
}

type GetLectureUseCase interface {
	GetLectureById(id int64) (*dto.GetLectureByIdOutputForm, error)
	GetLectures(limit int, offset int) (dto.GetLecturesOutputForm, error)

	GetParticipatingLecturesOfStudent(studentID int64, limit int, offset int) (dto.GetLecturesOutputForm, error)
	GetParticipatingLecturesOfAssistant(studentID int64, limit int, offset int) (dto.GetLecturesOutputForm, error)
	GetParticipatingLecturesOfTeacher(teacherID int64, limit int, offset int) (dto.GetLecturesOutputForm, error)
}

type CreateLectureUseCase interface {
	CreateLecture(form *dto.CreateLectureInputForm) (*dto.CreateLectureOutputForm, error)
}

type UpdateLectureUseCase interface {
	UpdateLecture(form *dto.UpdateLectureInputForm) (*dto.UpdateLectureOutputForm, error)
}

type DeleteLectureUseCase interface {
	DeleteLecture(id int64) (*dto.DeleteLectureOutputForm, error)
}
