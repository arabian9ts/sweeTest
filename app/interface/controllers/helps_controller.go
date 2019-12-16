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
	HelpsController struct {
		Student   *studentHelpsController
		Assistant *assistantHelpsController
		Teacher   *teacherHelpsController
	}

	studentHelpsController struct {
		InputPort port.HelpUseCase
		Validator validator.Validation
	}

	assistantHelpsController struct {
		InputPort port.HelpUseCase
		Validator validator.Validation
	}

	teacherHelpsController struct {
		InputPort port.HelpUseCase
		Validator validator.Validation
	}
)

func NewHelpsController(output port.HelpOutput, repository repository.HelpRepository, validator validator.Validation) (*HelpsController, error) {
	return &HelpsController{
		Student: &studentHelpsController{
			InputPort: &interactor.HelpInteractor{
				HelpRepository: repository,
				HelpOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &assistantHelpsController{
			InputPort: &interactor.HelpInteractor{
				HelpRepository: repository,
				HelpOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &teacherHelpsController{
			InputPort: &interactor.HelpInteractor{
				HelpRepository: repository,
				HelpOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

//
// for student helps
//
func (controller *studentHelpsController) GetHelpsByLectureId(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetHelpsByLectureID(lectureID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *studentHelpsController) CreateHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateHelpInputForm{StudentID: student.ID, LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateHelp(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *studentHelpsController) UpdateHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	helpID := getHelpID(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.UpdateHelpInputForm{ID: helpID, StudentID: student.ID, LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateHelp(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *studentHelpsController) DeleteHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	id := getHelpID(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteHelp(id, lectureID, student.ID)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

//
// for assistant helps
//
func (controller *assistantHelpsController) GetHelpsByLectureId(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetHelpsByLectureID(lectureID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *assistantHelpsController) DeleteHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	id := getHelpID(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteHelp(id, lectureID, student.ID)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

//
// for teacher helps
//
func (controller *teacherHelpsController) GetHelpsByLectureId(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetHelpsByLectureID(lectureID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *teacherHelpsController) DeleteHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	id := getHelpID(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteHelp(id, lectureID, student.ID)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
