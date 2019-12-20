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

type SubmissionsController struct {
	InputPort port.SubmissionUseCase
	Validator validator.Validation
}

func NewSubmissionsController(submissionRepository repository.SubmissionRepository, output port.SubmissionOutput, validator validator.Validation) (*SubmissionsController, error) {
	return &SubmissionsController{
		InputPort: &interactor.SubmissionInteractor{
			SubmissionRepository: submissionRepository,
			SubmissionOutput:     output,
		},
		Validator: validator,
	}, nil
}

func (controller *SubmissionsController) GetSubmissionTextsBySubmissionID(ctx Context) {
	submissionID, err := strconv.Atoi(ctx.Param("submission_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(ctx.Query("offset"))
	if err != nil {
		offset = 0
	}

	outputForm, err := controller.InputPort.GetSubmissionTextsBySubmissionID(int64(submissionID), limit, offset)
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *SubmissionsController) CreateSubmission(ctx Context) {
	taskID, err := strconv.Atoi(ctx.Param("task_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	inputForm := &dto.CreateSubmissionInputForm{TaskID: int64(taskID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateSubmission(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *SubmissionsController) CreateSubmissionText(ctx Context) {
	submissionID, err := strconv.Atoi(ctx.Param("submission_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}
	inputForm := &dto.CreateSubmissionTextInputForm{SubmissionID: int64(submissionID)}
	ctx.Bind(&inputForm)

	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.CreateSubmissionText(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *SubmissionsController) UpdateSubmissionText(ctx Context) {
	submissionID, err := strconv.Atoi(ctx.Param("submission_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	submissionTextID, err := strconv.Atoi(ctx.Param("submissionText_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	inputForm := &dto.UpdateSubmissionTextInputForm{ID: int64(submissionTextID), SubmissionID: int64(submissionID)}
	ok, msgs := controller.Validator.Validate(inputForm)
	if !ok {
		ctx.JSON(http.StatusBadRequest, msgs)
		return
	}

	outputForm, err := controller.InputPort.UpdateSubmissionText(inputForm)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}

func (controller *SubmissionsController) DeleteSubmissionText(ctx Context) {
	submissionID, err := strconv.Atoi(ctx.Param("submission_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	submissionTextID, err := strconv.Atoi(ctx.Param("submissionText_id"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, err)
		return
	}

	outputForm, err := controller.InputPort.DeleteSubmissionText(int64(submissionTextID), int64(submissionID))
	if err != nil {
		ctx.JSON(http.StatusServiceUnavailable, err)
		return
	}

	ctx.JSON(http.StatusOK, outputForm)
}
