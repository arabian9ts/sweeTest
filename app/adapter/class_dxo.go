package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateClassInputFormToClass(form *dto.CreateClassInputForm) (class *model.Class) {
	class = &model.Class{
		LectureID: form.LectureID,
		Name:      form.Name,
	}
	return
}

func ConvertUpdateClassInputFormToClass(form *dto.UpdateClassInputForm) (class *model.Class) {
	class = &model.Class{
		ID:        form.ID,
		LectureID: form.LectureID,
		Name:      form.Name,
	}
	return
}
