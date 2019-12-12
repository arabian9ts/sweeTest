package controllers

import (
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

func (controller *ClassesController) Index(ctx Context) {
	lectureID := getLectureID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetClassesByLectureId(int64(lectureID), limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ClassesController) Show(ctx Context) {
	classID, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	outputForm, err := controller.InputPort.GetClassById(int64(classID))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ClassesController) Create(ctx Context) {
	lectureID := getLectureID(ctx)
	inputForm := &dto.CreateClassInputForm{LectureID: int64(lectureID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateClass(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ClassesController) Update(ctx Context) {
	classID, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	lectureID := getLectureID(ctx)
	inputForm := &dto.UpdateClassInputForm{ID: int64(classID), LectureID: lectureID}
	ctx.Bind(&inputForm)
	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateClass(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *ClassesController) Delete(ctx Context) {
	id, err := getClassID(ctx)
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	lectureID := getLectureID(ctx)
	outputForm, err := controller.InputPort.DeleteClass(int64(id), int64(lectureID))
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}
