package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateTaskInputFormToTask(form *dto.CreateTaskInputForm) (*model.Task) {
	return &model.Task{
		LectureID: form.LectureID,
		Title:     form.Title,
		Desc:      form.Desc,
		Deadline:  form.Deadline,
	}
}

func ConvertUpdateTaskInputFormToTask(form *dto.UpdateTaskInputForm) (*model.Task) {
	return &model.Task{
		ID:        form.ID,
		LectureID: form.LectureID,
		Title:     form.Title,
		Desc:      form.Desc,
		Deadline:  form.Deadline,
	}
}
