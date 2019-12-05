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

func (*LecturePresenter) HandleGetLectures(lectures model.Lectures, err error) (dto.GetLecturesOutputForm, error) {
	output := dto.GetLecturesOutputForm{}
	for _, lecture := range lectures {
		form := &dto.GetLectureByIdOutputForm{
			ID:        lecture.ID,
			Name:      lecture.Name,
			CreatedAt: lecture.CreatedAt,
			UpdatedAt: lecture.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*LecturePresenter) HandleGetLectureById(lecture *model.Lecture, err error) (*dto.GetLectureByIdOutputForm, error) {
	output := &dto.GetLectureByIdOutputForm{
		ID:        lecture.ID,
		Name:      lecture.Name,
		CreatedAt: lecture.CreatedAt,
		UpdatedAt: lecture.UpdatedAt,
	}
	return output, err
}

func (*LecturePresenter) HandleCreateLecture(lecture *model.Lecture, err error) (*dto.CreateLectureOutputForm, error) {
	output := &dto.CreateLectureOutputForm{
		ID:        lecture.ID,
		Name:      lecture.Name,
		CreatedAt: lecture.CreatedAt,
		UpdatedAt: lecture.UpdatedAt,
	}
	return output, err
}

func (*LecturePresenter) HandleUpdateLecture(lecture *model.Lecture, err error) (*dto.UpdateLectureOutputForm, error) {
	output := &dto.UpdateLectureOutputForm{
		ID:        lecture.ID,
		Name:      lecture.Name,
		CreatedAt: lecture.CreatedAt,
		UpdatedAt: lecture.UpdatedAt,
	}
	return output, err
}

func (*LecturePresenter) HandleDeleteLecture(count int64, err error) (*dto.DeleteLectureOutputForm, error) {
	output := &dto.DeleteLectureOutputForm{AffectedRowsCount: count}
	return output, err
}
