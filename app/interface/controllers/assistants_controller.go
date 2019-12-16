package controllers

import (
	"net/http"
	"strconv"

	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type AssistantsController struct {
	InputPort port.UserUseCase
	Validator validator.Validation
}

func NewAssistantsController(userRepository repository.UserRepository, output port.UserOutput, validator validator.Validation) (*AssistantsController, error) {
	return &AssistantsController{
		InputPort: &interactor.UserInteractor{
			UserRepository: userRepository,
			UserOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *AssistantsController) GetAssistantById(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("assistant_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	outputForm, err := controller.InputPort.GetAssistantById(int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *AssistantsController) CreateAssistant(ctx Context) {
	inputForm := &dto.CreateAssistantInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateAssistant(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *AssistantsController) UpdateAssistant(ctx Context) {
	assistant := getCurrentAssistant(ctx)
	inputForm := &dto.UpdateAssistantInputForm{ID: assistant.ID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateAssistant(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
