package controllers

import (
	"net/http"
	"strconv"

	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type (
	ParticipantsController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}
)

func NewParticipantsController(repository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*ParticipantsController, error) {
	return &ParticipantsController{
		InputPort: &interactor.UserInteractor{
			UserRepository: repository,
			UserOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *ParticipantsController) GetStudentsByLectureId(ctx Context) {
	lectureId := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetStudentsByLectureId(lectureId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipantsController) GetAssistantsByLectureId(ctx Context) {
	lectureId := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetAssistantsByLectureId(lectureId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipantsController) GetTeachersByLectureId(ctx Context) {
	lectureId := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetTeachersByLectureId(lectureId, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
