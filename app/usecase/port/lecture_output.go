package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type LectureOutput interface {
	GetLectureOutput
	CreateLectureOutput
	UpdateLectureOutput
	DeleteLectureOutput
}

type GetLectureOutput interface {
	HandleGetLectures(total int64, lectures model.Lectures, err error) (dto.GetTotalLecturesOutputForm, error)
	HandleGetLectureById(lecture *model.Lecture, err error) (*dto.GetLectureByIdOutputForm, error)
}

type CreateLectureOutput interface {
	HandleCreateLecture(lecture *model.Lecture, err error) (*dto.CreateLectureOutputForm, error)
}

type UpdateLectureOutput interface {
	HandleUpdateLecture(lecture *model.Lecture, err error) (*dto.UpdateLectureOutputForm, error)
}

type DeleteLectureOutput interface {
	HandleDeleteLecture(count int64, err error) (*dto.DeleteLectureOutputForm, error)
}
