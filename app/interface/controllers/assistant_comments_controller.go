package controllers

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type AssistantCommentsController struct {
	InputPort port.CommentUseCase
	Validator validator.Validation
}

func NewAssistantCommentsController(output port.CommentOutput, repository repository.CommentRepository, validator validator.Validation) (*AssistantCommentsController, error) {
	return &AssistantCommentsController{
		InputPort: &interactor.CommentInteractor{
			CommentRepository: repository,
			Output:            output,
		},
		Validator: validator,
	}, nil
}

func (controller *AssistantCommentsController) Index(ctx Context) {
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

func (controller *AssistantCommentsController) Create(ctx Context) {
	assistant := getCurrentAssistant(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.CreateCommentInputForm{HelpID: helpID, UserType: model.AssistantUser, UserID: assistant.ID}
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

func (controller *AssistantCommentsController) Update(ctx Context) {
	id := getCommentID(ctx)
	assistant := getCurrentAssistant(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.UpdateCommentInputForm{ID: id, HelpID: helpID, UserType: model.AssistantUser, UserID: assistant.ID}
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

func (controller *AssistantCommentsController) Delete(ctx Context) {
	id := getCommentID(ctx)
	assistant := getCurrentAssistant(ctx)

	outputForm, err := controller.InputPort.DeleteComment(id, model.AssistantUser, assistant.ID)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
