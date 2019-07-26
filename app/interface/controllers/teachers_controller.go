package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
	"strconv"
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

func (controller *TeachersController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
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

	err := controller.Validator.Validate(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	teacher := adapter.ConvertTeacherInputFormToTeacher(inputForm)
	outputForm, err := controller.InputPort.CreateTeacher(teacher)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
