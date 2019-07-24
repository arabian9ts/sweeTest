package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type LecturePresenter struct {
}

func NewLecturePresenter() (port.LectureOutput) {
	return &LecturePresenter{}
}

func (*LecturePresenter) HandleGetLectures(lectures model.Lectures, err error) (dto.ReadLecturesOutputForm, error) {
	output := dto.ReadLecturesOutputForm{}
	for _, lecture := range lectures {
		form := &dto.ReadLectureOutputForm{
			ID:        lecture.ID,
			Name:      lecture.Name,
			CreatedAt: lecture.CreatedAt,
			UpdatedAt: lecture.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*LecturePresenter) HandleGetLectureById(lecture *model.Lecture, err error) (*dto.ReadLectureOutputForm, error) {
	output := &dto.ReadLectureOutputForm{
		ID:        lecture.ID,
		Name:      lecture.Name,
		CreatedAt: lecture.CreatedAt,
		UpdatedAt: lecture.UpdatedAt,
	}
	return output, err
}

func (*LecturePresenter) HandleCreateLecture(id int64, err error) (*dto.WriteLectureOutputForm, error) {
	output := &dto.WriteLectureOutputForm{LastChangedLectureId: id}
	return output, err
}

func (*LecturePresenter) HandleUpdateLecture(id int64, err error) (*dto.WriteLectureOutputForm, error) {
	output := &dto.WriteLectureOutputForm{LastChangedLectureId: id}
	return output, err
}

func (*LecturePresenter) HandleDeleteLecture(id int64, err error) (*dto.WriteLectureOutputForm, error) {
	output := &dto.WriteLectureOutputForm{LastChangedLectureId: id}
	return output, err
}
