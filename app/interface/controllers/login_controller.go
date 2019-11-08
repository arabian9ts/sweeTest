package controllers

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type (
	LoginController struct {
		Student   *studentLoginController
		Assistant *assistantLoginController
		Teacher   *teacherLoginController
		Admin     *adminLoginController
	}

	studentLoginController struct {
		InputPort port.AuthUseCase
		Validator validator.Validation
	}

	assistantLoginController struct {
		InputPort port.AuthUseCase
		Validator validator.Validation
	}

	teacherLoginController struct {
		InputPort port.AuthUseCase
		Validator validator.Validation
	}

	adminLoginController struct {
		InputPort port.AuthUseCase
		Validator validator.Validation
	}
)

func NewLoginController(userRepository repository.UserRepository, output port.AuthOutput, validator validator.Validation) (*LoginController, error) {
	return &LoginController{
		Student: &studentLoginController{
			InputPort: &interactor.AuthInteractor{
				UserRepository: userRepository,
				AuthOutput:     output,
			},
			Validator: validator,
		},
		Assistant: &assistantLoginController{
			InputPort: &interactor.AuthInteractor{
				UserRepository: userRepository,
				AuthOutput:     output,
			},
			Validator: validator,
		},
		Teacher: &teacherLoginController{
			InputPort: &interactor.AuthInteractor{
				UserRepository: userRepository,
				AuthOutput:     output,
			},
			Validator: validator,
		},
		Admin: &adminLoginController{
			InputPort: &interactor.AuthInteractor{
				UserRepository: userRepository,
				AuthOutput:     output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *studentLoginController) Create(ctx Context) {
	inputForm := &dto.AuthorizeStudentInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(401, msgs)
		return
	}

	outputForm, err := controller.InputPort.AuthorizeStudent(inputForm.StudentNo, inputForm.Password)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *assistantLoginController) Create(ctx Context) {
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

func (controller *teacherLoginController) Create(ctx Context) {
	inputForm := &dto.AuthorizeTeacherInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(401, msgs)
		return
	}

	outputForm, err := controller.InputPort.AuthorizeTeacher(inputForm.Email, inputForm.Password)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *adminLoginController) Create(ctx Context) {
	inputForm := &dto.AuthorizeAdminInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(401, msgs)
		return
	}

	outputForm, err := controller.InputPort.AuthorizeAdmin(inputForm.Email, inputForm.Password)
	if err != nil {
		ctx.JSON(401, err)
		return
	}

	ctx.JSON(200, outputForm)
}
