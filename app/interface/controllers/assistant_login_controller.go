package controllers

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type AssistantLoginController struct {
	InputPort port.AuthUseCase
	Validator validator.Validation
}

func NewAssistantLoginController(userRepository repository.UserRepository, output port.AuthOutput, validator validator.Validation) (*AssistantLoginController, error) {
	return &AssistantLoginController{
		InputPort: &interactor.AuthInteractor{
			UserRepository: userRepository,
			AuthOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *AssistantLoginController) Create(ctx Context) {
	inputForm := &dto.AuthorizeAssistantInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(401, msgs)
		return
	}

	outputForm, err := controller.InputPort.AuthorizeAssistant(inputForm.StudentNo, inputForm.Password)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	ctx.JSON(200, outputForm)
}
