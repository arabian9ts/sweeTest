package controllers

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
	"strconv"
)

type LecturesController struct {
	InputPort port.LectureUseCase
	Validator validator.Validation
}

func NewLecturesController(lectureRepository repository.LectureRepository, output port.LectureOutput, validator validator.Validation) (*LecturesController, error) {
	return &LecturesController{
		InputPort: &interactor.LectureInteractor{
			LectureRepository: lectureRepository,
			LectureOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *LecturesController) Index(ctx Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetLectures(limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *LecturesController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	outputForm, err := controller.InputPort.GetLectureById(int64(id))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *LecturesController) Create(ctx Context) {
	inputForm := &dto.CreateLectureInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateLecture(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *LecturesController) Update(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	inputForm := &dto.UpdateLectureInputForm{ID: int64(id)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateLecture(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *LecturesController) Delete(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("lecture_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	outputForm, err := controller.InputPort.DeleteLecture(int64(id))
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
