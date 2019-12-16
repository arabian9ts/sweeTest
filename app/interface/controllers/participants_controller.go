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
		Student   *studentParticipantsController
		Assistant *assistantParticipantsController
		Teacher   *teacherParticipantsController
	}

	// for student role
	studentParticipantsController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}

	// for assistant role
	assistantParticipantsController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}

	// for teacher role
	teacherParticipantsController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}
)

func NewParticipantsController(repository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*ParticipantsController, error) {
	return &ParticipantsController{
		Student: &studentParticipantsController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &assistantParticipantsController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &teacherParticipantsController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *studentParticipantsController) GetStudentsByLectureId(ctx Context) {
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

func (controller *assistantParticipantsController) GetStudentsByLectureId(ctx Context) {
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

func (controller *assistantParticipantsController) GetAssistantsByLectureId(ctx Context) {
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

func (controller *assistantParticipantsController) GetTeachersByLectureId(ctx Context) {
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

func (controller *teacherParticipantsController) GetStudentsByLectureId(ctx Context) {
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
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *teacherParticipantsController) GetAssistantsByLectureId(ctx Context) {
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
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *teacherParticipantsController) GetTeachersByLectureId(ctx Context) {
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
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
