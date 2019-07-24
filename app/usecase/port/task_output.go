package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type TaskOutput interface {
	GetTaskOutput
	CreateTaskOutput
	UpdateTaskOutput
	DeleteTaskOutput
}

type GetTaskOutput interface {
	HandleGetTasksByLectureId(tasks model.Tasks, err error) (dto.GetTasksByLectureIdOutputForm, error)
}

type CreateTaskOutput interface {
	HandleCreateTask(id int64, err error) (*dto.CreateTaskOutputForm, error)
}

type UpdateTaskOutput interface {
	HandleUpdateTask(count int64, err error) (*dto.UpdateTaskOutputForm, error)
}

type DeleteTaskOutput interface {
	HandleDeleteTask(count int64, err error) (*dto.DeleteTaskOutputForm, error)
}
