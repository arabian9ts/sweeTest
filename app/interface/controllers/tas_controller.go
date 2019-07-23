package controllers

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"strconv"
)

type TasController struct {
	InputPort port.UserUseCase
}

func NewTasController(userRepository repository.UserRepository, output port.UserOutput) (*TasController, error) {
	return &TasController{InputPort: &interactor.UserInteractor{
		UserRepository: userRepository,
		UserOutput:     output,
	}}, nil
}

func (controller *TasController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	outputForm, err := controller.InputPort.GetTaById(int64(id))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *TasController) Create(ctx Context) {
	inputForm := &dto.CreateTaInputForm{}
	ctx.Bind(&inputForm)
	ta := adapter.ConvertTaInputFormToTa(inputForm)

	outputForm, err := controller.InputPort.CreateTa(ta)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
