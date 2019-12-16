package controllers

import (
	"net/http"

	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type ClassesController struct {
	InputPort port.ClassUseCase
	Validator validator.Validation
}

func NewClassesController(classRepository repository.ClassRepository, output port.ClassOutput, validator validator.Validation) (*ClassesController, error) {
	return &ClassesController{
		InputPort: &interactor.ClassInteractor{
			ClassRepository: classRepository,
			ClassOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *ClassesController) GetClassesByLectureId(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetClassesByLectureId(lectureID, limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ClassesController) GetClassById(ctx Context) {
	classID, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	outputForm, err := controller.InputPort.GetClassById(classID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ClassesController) CreateClass(ctx Context) {
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateClassInputForm{LectureID: lectureID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateClass(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ClassesController) UpdateClass(ctx Context) {
	classID, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	lectureID := getLectureID(ctx)
	inputForm := &dto.UpdateClassInputForm{ID: classID, LectureID: lectureID}
	ctx.Bind(&inputForm)
	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateClass(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *ClassesController) DeleteClass(ctx Context) {
	classID, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	lectureID := getLectureID(ctx)
	outputForm, err := controller.InputPort.DeleteClass(classID, lectureID)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
