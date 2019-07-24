package interactor

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
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

func (interactor *TaskInteractor) GetTasksByLectureId(lectureId int64, limit int, offset int) (dto.GetTasksByLectureIdOutputForm, error) {
	return interactor.TaskOutput.HandleGetTasksByLectureId(
		interactor.TaskRepository.GetTasksByLectureId(lectureId, limit, offset),
	)
}

func (interactor *TaskInteractor) CreateTask(task *model.Task) (*dto.CreateTaskOutputForm, error) {
	return interactor.TaskOutput.HandleCreateTask(
		interactor.TaskRepository.CreateTask(task),
	)
}

func (interactor *TaskInteractor) UpdateTask(task *model.Task) (*dto.UpdateTaskOutputForm, error) {
	return interactor.TaskOutput.HandleUpdateTask(
		interactor.TaskRepository.UpdateTask(task),
	)
}

func (interactor *TaskInteractor) DeleteTask(id int64, lectureID int64) (*dto.DeleteTaskOutputForm, error) {
	return interactor.TaskOutput.HandleDeleteTask(
		interactor.TaskRepository.DeleteTask(id, lectureID),
	)
}
