package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type SubmissionRepository interface {
	GetSubmissionsByStudentID(studentID int64, limit int, offset int) (submissions model.Submissions, err error)
	GetSubmissionsByTaskID(taskID int64, limit int, offset int) (submissions model.Submissions, err error)
	GetSubmissionTextByID(submissionTextID int64) (submissionText *model.SubmissionText, err error)
	GetSubmissionTextsBySubmissionID(submissionID int64, limit int, offset int) (submissionTexts model.SubmissionTexts, err error)

	CreateSubmissionText(submission *model.Submission, submissionText *model.SubmissionText) (*model.Submission, *model.SubmissionText, error)
	UpdateSubmissionText(submissionText *model.SubmissionText) (*model.SubmissionText, error)
	DeleteSubmissionText(id int64, submissionID int64) (int64, error)
}
