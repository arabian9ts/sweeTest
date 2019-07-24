package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
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
}

type CreateLectureUseCase interface {
	CreateLecture(lecture *model.Lecture) (*dto.CreateLectureOutputForm, error)
}

type UpdateLectureUseCase interface {
	UpdateLecture(lecture *model.Lecture) (*dto.UpdateLectureOutputForm, error)
}

type DeleteLectureUseCase interface {
	DeleteLecture(id int64) (*dto.DeleteLectureOutputForm, error)
}
