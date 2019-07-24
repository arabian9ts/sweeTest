package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type TaskUseCase interface {
	GetTaskUseCase
	CreateTaskUseCase
	UpdateTaskUseCase
	DeleteTaskUseCase
}

type GetTaskUseCase interface {
	GetTasksByLectureId(lectureId int64, limit int, offset int) (dto.GetTasksByLectureIdOutputForm, error)
}

type CreateTaskUseCase interface {
	CreateTask(task *model.Task) (*dto.CreateTaskOutputForm, error)
}

type UpdateTaskUseCase interface {
	UpdateTask(task *model.Task) (*dto.UpdateTaskOutputForm, error)
}

type DeleteTaskUseCase interface {
	DeleteTask(id int64, lectureID int64) (*dto.DeleteTaskOutputForm, error)
}
