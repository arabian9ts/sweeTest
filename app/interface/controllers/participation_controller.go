package controllers

import (
	"net/http"

	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type (
	ParticipationController struct {
		InputPort port.ParticipationUseCase
		Validator validator.Validation
	}
)

func NewParticipationController(repository repository.ParticipationRepository, output port.ParticipationOutput, validator validator.Validation) (*ParticipationController, error) {
	return &ParticipationController{
		InputPort: &interactor.ParticipationInteractor{
			ParticipationRepository: repository,
			ParticipationOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *ParticipationController) StudentParticipateToLecture(ctx Context) {
	student := getCurrentStudent(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateStudentLectureInputForm{StudentID: student.ID, LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateStudentLecture(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipationController) StudentExitFromLecture(ctx Context) {
	student := getCurrentStudent(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteStudentLecture(student.ID, lectureID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipationController) AssistantParticipateToLecture(ctx Context) {
	assistant := getCurrentAssistant(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateAssistantLectureInputForm{StudentID: assistant.ID, LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateAssistantLecture(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipationController) AssistantExitFromLecture(ctx Context) {
	assistant := getCurrentAssistant(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteAssistantLecture(assistant.ID, lectureID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipationController) TeacherParticipateToLecture(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateTeacherLectureInputForm{TeacherID: teacher.ID, LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateTeacherLecture(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ParticipationController) TeacherExitFromLecture(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteTeacherLecture(teacher.ID, lectureID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
