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

type StudentsController struct {
	InputPort port.UserUseCase
	Validator validator.Validation
}

func NewStudentsController(userRepository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*StudentsController, error) {
	return &StudentsController{
		InputPort: &interactor.UserInteractor{
			UserRepository: userRepository,
			UserOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *StudentsController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	outputForm, err := controller.InputPort.GetStudentById(int64(id))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *StudentsController) Create(ctx Context) {
	inputForm := &dto.CreateStudentInputForm{}
	ctx.Bind(&inputForm)

	err := controller.Validator.Validate(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	student := adapter.ConvertStudentInputFormToUser(inputForm)
	outputForm, err := controller.InputPort.CreateStudent(student)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
