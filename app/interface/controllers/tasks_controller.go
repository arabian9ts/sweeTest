package controllers

import (
	"net/http"
	"strconv"

	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type TasksController struct {
	InputPort port.TaskUseCase
	Validator validator.Validation
}

func NewTasksController(taskRepository repository.TaskRepository, output port.TaskOutput, validator validator.Validation) (*TasksController, error) {
	return &TasksController{
		InputPort: &interactor.TaskInteractor{
			TaskRepository: taskRepository,
			TaskOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *TasksController) GetTasksByLectureId(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetTasksByLectureId(int64(lectureID), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *TasksController) CreateTask(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	inputForm := &dto.CreateTaskInputForm{LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateTask(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *TasksController) UpdateTask(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	id, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	inputForm := &dto.UpdateTaskInputForm{ID: int64(id), LectureID: int64(lectureID)}
	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateTask(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *TasksController) DeleteTask(ctx Context) {
	lectureID, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	taskID, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	outputForm, err := controller.InputPort.DeleteTask(int64(taskID), int64(lectureID))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
