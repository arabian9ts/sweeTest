package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateSubmissionInputFormToSubmission(form *dto.CreateSubmissionInputForm) *model.Submission {
	return &model.Submission{
		TaskID:    form.TaskID,
		StudentID: form.StudentID,
	}
}

func ConvertCreateSubmissionTextInputFormToSubmissionText(form *dto.CreateSubmissionTextInputForm) *model.SubmissionText {
	return &model.SubmissionText{
		SubmissionID: form.SubmissionID,
		FileName:     form.FileName,
		Contents:     form.Contents,
	}
}

func ConvertUpdateSubmissionTextInputFormToSubmissionText(form *dto.UpdateSubmissionTextInputForm) *model.SubmissionText {
	return &model.SubmissionText{
		ID:           form.ID,
		SubmissionID: form.SubmissionID,
		FileName:     form.FileName,
		Contents:     form.Contents,
	}
}
