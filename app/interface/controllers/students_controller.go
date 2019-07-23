package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type StudentsController struct {
	InputPort port.UserUseCase
}

func NewUsersController(userRepository repository.UserRepository, output port.UserOutput) (*StudentsController, error) {
	return &StudentsController{InputPort: &interactor.UserInteractor{
		UserRepository: userRepository,
		UserOutput:     output,
	}}, nil
}

func (controller *StudentsController) Create(ctx Context) {
	inputForm := &dto.CreateStudentInputForm{}
	ctx.Bind(&inputForm)
	student := adapter.ConvertStudentInputFormToUser(inputForm)

	outputForm, err := controller.InputPort.CreateStudent(student)
	if err != nil {
		ctx.JSON(400, err)
	}

	ctx.JSON(200, outputForm)
}
