package controllers

import (
	"strconv"

	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type (
	ParticipatingLecturesController struct {
		Student   *studentLectureController
		Assistant *assistantLectureController
		Teacher   *teacherLectureController
	}

	studentLectureController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}

	assistantLectureController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}

	teacherLectureController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}
)

func NewParticipatingLecturesController(repository repository.LectureRepository, output port.LectureOutput, validator validator.Validation) (*ParticipatingLecturesController, error) {
	return &ParticipatingLecturesController{
		Student: &studentLectureController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &assistantLectureController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &teacherLectureController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *studentLectureController) Index(ctx Context) {
	student := getCurrentStudent(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatingLecturesOfStudent(student.ID, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *assistantLectureController) Index(ctx Context) {
	assistant := getCurrentAssistant(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatingLecturesOfAssistant(assistant.ID, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *teacherLectureController) Index(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatingLecturesOfTeacher(teacher.ID, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
