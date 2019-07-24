package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertLectureInputFormToLecture(form *dto.CreateLectureInputForm) (lecture *model.Lecture) {
	lecture = &model.Lecture{
		Name: form.Name,
	}
	return
}
