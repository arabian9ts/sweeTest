package controllers

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type StudentCommentsController struct {
	InputPort port.CommentUseCase
	Validator validator.Validation
}

func NewStudentCommentsController(output port.CommentOutput, repository repository.CommentRepository, validator validator.Validation) (*StudentCommentsController, error) {
	return &StudentCommentsController{
		InputPort: &interactor.CommentInteractor{
			CommentRepository: repository,
			Output:            output,
		},
		Validator: validator,
	}, nil
}

func (controller *StudentCommentsController) Index(ctx Context) {
	_ = getCurrentStudent(ctx)
	helpID := getHelpID(ctx)
	limit := getLimit(ctx)
	offset := getOffset(ctx)

	outputForm, err := controller.InputPort.GetCommentsByHelpID(helpID, limit, offset)
	if err != nil {
		ctx.JSON(503, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *StudentCommentsController) Create(ctx Context) {
	student := getCurrentStudent(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.CreateCommentInputForm{HelpID: helpID, UserType: model.StudentUser, UserID: student.ID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateComment(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *StudentCommentsController) Update(ctx Context) {
	id := getCommentID(ctx)
	student := getCurrentStudent(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.UpdateCommentInputForm{ID: id, HelpID: helpID, UserType: model.StudentUser, UserID: student.ID}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(400, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateComment(inputForm)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *StudentCommentsController) Delete(ctx Context) {
	id := getCommentID(ctx)
	student := getCurrentStudent(ctx)

	outputForm, err := controller.InputPort.DeleteComment(id, model.StudentUser, student.ID)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}