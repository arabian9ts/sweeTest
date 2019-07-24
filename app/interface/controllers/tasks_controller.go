package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"strconv"
)

type TasksController struct {
	InputPort port.TaskUseCase
}

func NewTasksController(taskRepository repository.TaskRepository, output port.TaskOutput) (*TasksController, error) {
	return &TasksController{InputPort: &interactor.TaskInteractor{
		TaskRepository: taskRepository,
		TaskOutput:     output,
	}}, nil
}

func (controller *TasksController) Index(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetTasksByLectureId(int64(lectureID), limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TasksController) Create(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	inputForm := &dto.CreateTaskInputForm{LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)
	task := adapter.ConvertCreateTaskInputFormToTask(inputForm)

	outputForm, err := controller.InputPort.CreateTask(task)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TasksController) Update(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	id, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	inputForm := &dto.UpdateTaskInputForm{ID: int64(id), LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)
	task := adapter.ConvertUpdateTaskInputFormToTask(inputForm)

	outputForm, err := controller.InputPort.UpdateTask(task)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TasksController) Delete(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	taskID, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	outputForm, err := controller.InputPort.DeleteTask(int64(taskID), int64(lectureID))
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
