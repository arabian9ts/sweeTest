package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type LectureOutput interface {
	LectureReadOutput
	LectureWriteOutput
}

type LectureReadOutput interface {
	HandleGetLectures(lectures model.Lectures, err error) (dto.ReadLecturesOutputForm, error)
	HandleGetLectureById(lecture *model.Lecture, err error) (*dto.ReadLectureOutputForm, error)
}

type LectureWriteOutput interface {
	HandleCreateLecture(id int64, err error) (*dto.WriteLectureOutputForm, error)
	HandleUpdateLecture(id int64, err error) (*dto.WriteLectureOutputForm, error)
	HandleDeleteLecture(id int64, err error) (*dto.WriteLectureOutputForm, error)
}
