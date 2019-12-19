package port

import "github.com/arabian9ts/sweeTest/app/dto"

type SubmissionUseCase interface {
	GetSubmissionUseCase
	CreateSubmissionUseCase
	UpdateSubmissionUseCase
	DeleteSubmissionUseCase
}

type GetSubmissionUseCase interface {
	GetSubmissionTextByID(submissionTextID int64) (*dto.GetSubmissionTextOutputForm, error)
	GetSubmissionsTextBySubmissionID(submissionID int64, limit int, offset int) (dto.GetSubmissionTextsOutputForm, error)
}

type CreateSubmissionUseCase interface {
	CreateSubmission(form *dto.CreateSubmissionInputForm) (*dto.CreateSubmissionOutputForm, error)
	CreateSubmissionText(form *dto.CreateSubmissionTextInputForm) (*dto.CreateSubmissionTextOutputForm, error)
}

type UpdateSubmissionUseCase interface {
	UpdateSubmissionText(form *dto.UpdateSubmissionTextInputForm) (*dto.UpdateSubmissionTextOutputForm, error)
}

type DeleteSubmissionUseCase interface {
	DeleteSubmission(id int64, submissionID int64) (*dto.DeleteSubmissionTextOutputForm, error)
}
