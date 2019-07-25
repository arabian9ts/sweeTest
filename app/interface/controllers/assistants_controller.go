package controllers

import (
	"strconv"

	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type AssistantsController struct {
	InputPort port.UserUseCase
}

func NewAssistantsController(userRepository repository.UserRepository, output port.UserOutput) (*AssistantsController, error) {
	return &AssistantsController{InputPort: &interactor.UserInteractor{
		UserRepository: userRepository,
		UserOutput:     output,
	}}, nil
}

func (controller *AssistantsController) Show(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(404, err)
		return
	}
	outputForm, err := controller.InputPort.GetAssistantById(int64(id))
	if err != nil {
		ctx.JSON(404, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *AssistantsController) Create(ctx Context) {
	inputForm := &dto.CreateAssistantInputForm{}
	ctx.Bind(&inputForm)
	assistant := adapter.ConvertAssistantInputFormToAssistant(inputForm)

	outputForm, err := controller.InputPort.CreateAssistant(assistant)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
