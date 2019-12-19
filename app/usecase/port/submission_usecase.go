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
	GetSubmissionsByStudentID(StudentID int64, limit int, offset int) (dto.GetSubmissionsOutputForm, error)
	GetSubmissionsByTaskID(taskID int64, limit int, offset int) (dto.GetSubmissionsOutputForm, error)
}

type CreateSubmissionUseCase interface {
	CreateSubmissionText(form *dto.CreateSubmissionTextInputForm) (*dto.CreateSubmissionTextOutputForm, error)
}

type UpdateSubmissionUseCase interface {
	UpdateSubmissionText(form *dto.UpdateSubmissionTextInputForm) (*dto.UpdateSubmissionTextOutputForm, error)
}

type DeleteSubmissionUseCase interface {
	DeleteSubmission(id int64, submissionID int64) (*dto.DeleteSubmissionTextOutputForm, error)
}
