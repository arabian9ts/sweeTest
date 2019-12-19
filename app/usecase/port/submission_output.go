package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type SubmissionOutput interface {
	GetSubmissionOutput
	CreateSubmissionOutput
	UpdateSubmissionOutput
	DeleteSubmissionOutput
}

type GetSubmissionOutput interface {
	HandleGetSubmissionTextsBySubmissionID(submissionTexts model.SubmissionTexts, err error) (dto.GetSubmissionTextsOutputForm, error)
	HandleGetSubmissionTextByID(submissionText *model.SubmissionText, err error) (*dto.GetSubmissionTextOutputForm, error)
}

type CreateSubmissionOutput interface {
	HandleCreateSubmission(submission *model.Submission, err error) (*dto.CreateSubmissionOutputForm, error)
	HandleCreateSubmissionText(submissionText *model.SubmissionText, err error) (*dto.CreateSubmissionTextOutputForm, error)
}

type UpdateSubmissionOutput interface {
	HandleUpdateSubmissionText(submissionText *model.SubmissionText, err error) (*dto.UpdateSubmissionTextOutputForm, error)
}

type DeleteSubmissionOutput interface {
	HandleDeleteSubmissionText(id int64, err error) (*dto.DeleteSubmissionTextOutputForm, error)
}