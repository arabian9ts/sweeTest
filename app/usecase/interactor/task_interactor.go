package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type TaskInteractor struct {
	TaskRepository repository.TaskRepository
	TaskOutput     port.TaskOutput
}

func NewTaskInteractor(repository repository.TaskRepository, output port.TaskOutput) (*TaskInteractor, error) {
	return &TaskInteractor{TaskRepository: repository, TaskOutput: output}, nil
}

func (interactor *TaskInteractor) GetTasksByClassId(classId int64, limit int, offset int) (dto.GetTasksByClassIdOutputForm, error) {
	return interactor.TaskOutput.HandleGetTasksByClassId(
		interactor.TaskRepository.GetTasksByClassId(classId, limit, offset),
	)
}

func (interactor *TaskInteractor) CreateTask(form *dto.CreateTaskInputForm) (*dto.CreateTaskOutputForm, error) {
	task := adapter.ConvertCreateTaskInputFormToTask(form)
	return interactor.TaskOutput.HandleCreateTask(
		interactor.TaskRepository.CreateTask(task),
	)
}

func (interactor *TaskInteractor) UpdateTask(form *dto.UpdateTaskInputForm) (*dto.UpdateTaskOutputForm, error) {
	task := adapter.ConvertUpdateTaskInputFormToTask(form)
	return interactor.TaskOutput.HandleUpdateTask(
		interactor.TaskRepository.UpdateTask(task),
	)
}

func (interactor *TaskInteractor) DeleteTask(id int64, classID int64) (*dto.DeleteTaskOutputForm, error) {
	return interactor.TaskOutput.HandleDeleteTask(
		interactor.TaskRepository.DeleteTask(id, classID),
	)
}
