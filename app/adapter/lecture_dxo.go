package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateLectureInputFormToLecture(form *dto.CreateLectureInputForm) (lecture *model.Lecture) {
	lecture = &model.Lecture{
		Name: form.Name,
	}
	return
}

func ConvertUpdateLectureInputFormToLecture(form *dto.UpdateLectureInputForm) (lecture *model.Lecture) {
	lecture = &model.Lecture{
		ID:   form.ID,
		Name: form.Name,
	}
	return
}
