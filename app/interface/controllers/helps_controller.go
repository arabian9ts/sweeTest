package controllers

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type HelpsController struct {
	InputPort port.HelpUseCase
	Validator validator.Validation
}

func NewHelpsController(output port.HelpOutput, repository repository.HelpRepository, validator validator.Validation) (*HelpsController, error) {
	return &HelpsController{
		InputPort: &interactor.HelpInteractor{
			HelpRepository: repository,
			HelpOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *HelpsController) GetHelpsByLectureId(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetHelpsByLectureID(int64(lectureID), limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *HelpsController) CreateHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateHelpInputForm{StudentID: student.ID, LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateHelp(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *HelpsController) UpdateHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	helpID := getHelpID(ctx)
	lectureID := getLectureID(ctx)
	inputForm := &dto.UpdateHelpInputForm{ID: int64(helpID), StudentID: student.ID, LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateHelp(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *HelpsController) DeleteHelp(ctx Context) {
	student := getCurrentStudent(ctx)
	id := getHelpID(ctx)
	lectureID := getLectureID(ctx)

	outputForm, err := controller.InputPort.DeleteHelp(int64(id), int64(lectureID), student.ID)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
