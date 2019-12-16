package controllers

import (
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
	lecture := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetStudentsByLectureId(lecture, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ParticipantsController) GetAssistantsByLectureId(ctx Context) {
	lecture := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetAssistantsByLectureId(lecture, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ParticipantsController) GetTeachersByLectureId(ctx Context) {
	lecture := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetTeachersByLectureId(lecture, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
