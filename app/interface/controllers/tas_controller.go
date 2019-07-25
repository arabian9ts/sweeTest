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

type TasController struct {
	InputPort port.UserUseCase
	Validator validator.Validation
}

func NewTasController(userRepository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*TasController, error) {
	return &TasController{
		InputPort: &interactor.UserInteractor{
			UserRepository: userRepository,
			UserOutput:     output,
		},
		Validator: validator,
	}, nil
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

	err := controller.Validator.Validate(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ta := adapter.ConvertTaInputFormToTa(inputForm)
	outputForm, err := controller.InputPort.CreateTa(ta)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
