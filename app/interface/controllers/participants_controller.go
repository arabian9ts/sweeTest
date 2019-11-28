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
		Student   *participatedStudentController
		Assistant *participatedAssistantController
		Teacher   *participatedTeacherController
	}

	participatedStudentController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}

	participatedAssistantController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}

	participatedTeacherController struct {
		InputPort port.UserUseCase
		Validator validator.Validation
	}
)

func NewParticipantsController(repository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*ParticipantsController, error) {
	return &ParticipantsController{
		Student: &participatedStudentController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &participatedAssistantController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &participatedTeacherController{
			InputPort: &interactor.UserInteractor{
				UserRepository: repository,
				UserOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *participatedStudentController) Index(ctx Context) {
	lecture := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatedStudents(lecture, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *participatedAssistantController) Index(ctx Context) {
	lecture := getLectureID(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatedAssistants(lecture, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *participatedTeacherController) Index(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetParticipatedTeachers(teacher.ID, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
