package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"strconv"
)

type TeachersController struct {
	InputPort port.UserUseCase
}

func NewTeachersController(userRepository repository.UserRepository, output port.UserOutput) (*TeachersController, error) {
	return &TeachersController{InputPort: &interactor.UserInteractor{
		UserRepository: userRepository,
		UserOutput:     output,
	}}, nil
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
	teacher := adapter.ConvertTeacherInputFormToTeacher(inputForm)

	outputForm, err := controller.InputPort.CreateTeacher(teacher)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
