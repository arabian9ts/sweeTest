package port

import "github.com/arabian9ts/sweeTest/app/dto"

type SubmissionUseCase interface {
	GetSubmissionUseCase
	CreateSubmissionUseCase
	UpdateSubmissionUseCase
	DeleteSubmissionUseCase
}

type GetSubmissionUseCase interface {
	GetSubmissionByID(submissionTextID int64) (*dto.GetSubmissionOutputForm, error)
	GetSubmissionsByTaskID(taskID int64, limit int, offset int) (dto.GetSubmissionsOutputForm, error)
}

type CreateSubmissionUseCase interface {
	CreateSubmission(form *dto.CreateSubmissionInputForm) (*dto.CreateSubmissionOutputForm, error)
}

type UpdateSubmissionUseCase interface {
	UpdateSubmission(form *dto.UpdateSubmissionInputForm) (*dto.UpdateSubmissionOutputForm, error)
}

type DeleteSubmissionUseCase interface {
	DeleteSubmission(id int64, taskID int64) (*dto.DeleteSubmissionOutputForm, error)
}
