package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateTaskInputFormToTask(form *dto.CreateTaskInputForm) *model.Task {
	return &model.Task{
		ClassID:  form.ClassID,
		Title:    form.Title,
		Desc:     form.Desc,
		Deadline: form.Deadline,
	}
}

func ConvertUpdateTaskInputFormToTask(form *dto.UpdateTaskInputForm) *model.Task {
	return &model.Task{
		ID:       form.ID,
		ClassID:  form.ClassID,
		Title:    form.Title,
		Desc:     form.Desc,
		Deadline: form.Deadline,
	}
}
