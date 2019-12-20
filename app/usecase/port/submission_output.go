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
	HandleGetSubmissionsByTaskID(submissionTexts model.Submissions, err error) (dto.GetSubmissionsOutputForm, error)
	HandleGetSubmissionByID(submission model.Submission, err error) (*dto.GetSubmissionOutputForm, error)
}

type CreateSubmissionOutput interface {
	HandleCreateSubmission(submission *model.Submission, err error) (*dto.CreateSubmissionOutputForm, error)
}

type UpdateSubmissionOutput interface {
	HandleUpdateSubmission(submission *model.Submission, err error) (*dto.UpdateSubmissionOutputForm, error)
}

type DeleteSubmissionOutput interface {
	HandleDeleteSubmission(id int64, taskID int64, err error) (*dto.DeleteSubmissionTextOutputForm, error)
}
