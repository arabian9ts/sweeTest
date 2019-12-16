package controllers

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
)

type (
	CommentsController struct {
		Student   *studentCommentsController
		Assistant *assistantCommentsController
		Teacher   *teacherCommentsController
	}

	studentCommentsController struct {
		InputPort port.CommentUseCase
		Validator validator.Validation
	}

	assistantCommentsController struct {
		InputPort port.CommentUseCase
		Validator validator.Validation
	}

	teacherCommentsController struct {
		InputPort port.CommentUseCase
		Validator validator.Validation
	}
)

func NewCommentsController(output port.CommentOutput, repository repository.CommentRepository, validator validator.Validation) (*CommentsController, error) {
	return &CommentsController{
		Student: &studentCommentsController{
			InputPort: &interactor.CommentInteractor{
				CommentRepository: repository,
				Output:            output,
			},
			Validator: validator,
		},
		Assistant: &assistantCommentsController{
			InputPort: &interactor.CommentInteractor{
				CommentRepository: repository,
				Output:            output,
			},
			Validator: validator,
		},
		Teacher: &teacherCommentsController{
			InputPort: &interactor.CommentInteractor{
				CommentRepository: repository,
				Output:            output,
			},
			Validator: validator,
		},
	}, nil
}

func (controller *studentCommentsController) GetCommentsByHelpId(ctx Context) {
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

func (controller *studentCommentsController) CreateComment(ctx Context) {
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

func (controller *studentCommentsController) UpdateComment(ctx Context) {
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

func (controller *studentCommentsController) DeleteComment(ctx Context) {
	id := getCommentID(ctx)
	student := getCurrentStudent(ctx)

	outputForm, err := controller.InputPort.DeleteComment(id, model.StudentUser, student.ID)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *assistantCommentsController) GetCommentsByHelpId(ctx Context) {
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

func (controller *assistantCommentsController) CreateComment(ctx Context) {
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

func (controller *assistantCommentsController) UpdateComment(ctx Context) {
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

func (controller *assistantCommentsController) DeleteComment(ctx Context) {
	id := getCommentID(ctx)
	assistant := getCurrentAssistant(ctx)

	outputForm, err := controller.InputPort.DeleteComment(id, model.AssistantUser, assistant.ID)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}

func (controller *teacherCommentsController) Index(ctx Context) {
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

func (controller *teacherCommentsController) Create(ctx Context) {
	teacher := getCurrentTeacher(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.CreateCommentInputForm{HelpID: helpID, UserType: model.TeacherUser, UserID: teacher.ID}
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

func (controller *teacherCommentsController) Update(ctx Context) {
	id := getCommentID(ctx)
	teacher := getCurrentTeacher(ctx)
	helpID := getHelpID(ctx)
	inputForm := &dto.UpdateCommentInputForm{ID: id, HelpID: helpID, UserType: model.TeacherUser, UserID: teacher.ID}
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

func (controller *teacherCommentsController) Delete(ctx Context) {
	id := getCommentID(ctx)
	teacher := getCurrentTeacher(ctx)

	outputForm, err := controller.InputPort.DeleteComment(id, model.TeacherUser, teacher.ID)
	if err != nil {
		ctx.JSON(400, err)
		return
	}

	ctx.JSON(200, outputForm)
}
