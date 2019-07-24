package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type LectureUseCase interface {
	LectureReadUseCase
	LectureWriteUseCase
}

type LectureReadUseCase interface {
	GetLectureById(id int64) (*dto.ReadLectureOutputForm, error)
	GetLectures(limit int, offset int) (dto.ReadLecturesOutputForm, error)
}

type LectureWriteUseCase interface {
	CreateLecture(lecture *model.Lecture) (*dto.WriteLectureOutputForm, error)
	UpdateLecture(lecture *model.Lecture) (*dto.WriteLectureOutputForm, error)
	DeleteLecture(id int64) (*dto.WriteLectureOutputForm, error)
}
