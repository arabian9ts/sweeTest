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
	ParticipatingLecturesController struct {
		Student   *studentParticipatingLecturesController
		Assistant *assistantParticipatingLecturesController
		Teacher   *teacherParticipatingLecturesController
	}

	studentParticipatingLecturesController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}

	assistantParticipatingLecturesController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}

	teacherParticipatingLecturesController struct {
		InputPort port.LectureUseCase
		Validator validator.Validation
	}
)

func NewParticipatingLecturesController(repository repository.LectureRepository, output port.LectureOutput, validator validator.Validation) (*ParticipatingLecturesController, error) {
	return &ParticipatingLecturesController{
		Student: &studentParticipatingLecturesController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &assistantParticipatingLecturesController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &teacherParticipatingLecturesController{
			InputPort: &interactor.LectureInteractor{
				LectureRepository: repository,
				LectureOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *studentParticipatingLecturesController) GetLectures(ctx Context) {
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
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *assistantParticipatingLecturesController) GetLectures(ctx Context) {
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
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *teacherParticipatingLecturesController) GetLectures(ctx Context) {
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
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
