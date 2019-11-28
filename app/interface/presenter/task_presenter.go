package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type TaskPresenter struct{}

func NewTaskPresenter() port.TaskOutput {
	return &TaskPresenter{}
}

func (*TaskPresenter) HandleGetTasksByLectureId(tasks model.Tasks, err error) (dto.GetTasksByLectureIdOutputForm, error) {
	output := dto.GetTasksByLectureIdOutputForm{}
	for _, task := range tasks {
		form := &dto.GetTaskOutputForm{
			ID:        task.ID,
			LectureID: task.LectureID,
			Title:     task.Title,
			Desc:      task.Desc,
			Deadline:  task.Deadline,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		}
		output = append(output, form)
	}

	return output, err
}

func (*TaskPresenter) HandleCreateTask(task *model.Task, err error) (*dto.CreateTaskOutputForm, error) {
	output := &dto.CreateTaskOutputForm{
		ID:        task.ID,
		LectureID: task.LectureID,
		Title:     task.Title,
		Desc:      task.Desc,
		Deadline:  task.Deadline,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return output, err
}

func (*TaskPresenter) HandleUpdateTask(task *model.Task, err error) (*dto.UpdateTaskOutputForm, error) {
	output := &dto.UpdateTaskOutputForm{
		ID:        task.ID,
		LectureID: task.LectureID,
		Title:     task.Title,
		Desc:      task.Desc,
		Deadline:  task.Deadline,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return output, err
}

func (*TaskPresenter) HandleDeleteTask(count int64, err error) (*dto.DeleteTaskOutputForm, error) {
	output := &dto.DeleteTaskOutputForm{AffectedRowsCount: count}

	return output, err
}
