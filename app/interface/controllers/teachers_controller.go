package controllers

import (
	"strconv"

	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type TeachersController struct {
	InputPort port.UserUseCase
	Validator validator.Validation
}

func NewTeachersController(userRepository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*TeachersController, error) {
	return &TeachersController{
		InputPort: &interactor.UserInteractor{
			UserRepository: userRepository,
			UserOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *TeachersController) Index(ctx Context) {
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetTeachers(limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TeachersController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("teacher_id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	outputForm, err := controller.InputPort.GetTeacherById(int64(id))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TeachersController) Create(ctx Context) {
	inputForm := &dto.CreateTeacherInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateTeacher(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TeachersController) Update(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	inputForm := &dto.UpdateTeacherInputForm{ID: teacher.ID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateTeacher(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
