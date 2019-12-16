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

func (controller *StudentsController) GetStudents(ctx Context) {
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetStudents(limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *StudentsController) GetStudentById(ctx Context) {
	id, err := strconv.Atoi(ctx.Param("student_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	outputForm, err := controller.InputPort.GetStudentById(int64(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *StudentsController) SignUp(ctx Context) {
	inputForm := &dto.CreateStudentInputForm{}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateStudent(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *StudentsController) UpdateStudent(ctx Context) {
	student := getCurrentStudent(ctx)
	inputForm := &dto.UpdateStudentInputForm{ID: student.ID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateStudent(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
