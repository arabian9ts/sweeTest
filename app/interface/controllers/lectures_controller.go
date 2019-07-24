package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"strconv"
)

type LecturesController struct {
	InputPort port.LectureUseCase
}

func NewLecturesController(lectureRepository repository.LectureRepository, output port.LectureOutput) (*LecturesController, error) {
	return &LecturesController{InputPort: &interactor.LectureInteractor{
		LectureRepository: lectureRepository,
		LectureOutput:     output,
	}}, nil
}

func (controller *LecturesController) Index(ctx Context) {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Param("offset"))
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
	lecture := adapter.ConvertLectureInputFormToLecture(inputForm)

	outputForm, err := controller.InputPort.CreateLecture(lecture)
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

	inputForm := &dto.CreateLectureInputForm{}
	ctx.Bind(&inputForm)
	lecture := adapter.ConvertLectureInputFormToLecture(inputForm)
	lecture.ID = int64(id)

	outputForm, err := controller.InputPort.UpdateLecture(lecture)
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
