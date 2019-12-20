package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type SubmissionRepository interface {
	GetSubmissionByID(submissionID int64) (submission *model.Submission, err error)
	GetSubmissionsByTaskID(taskID int64, limit int, offset int) (submissions model.Submissions, err error)
	CreateSubmission(submission *model.Submission) (*model.Submission, error)
	UpdateSubmission(submissionText *model.Submission) (*model.Submission, error)
	DeleteSubmission(id int64, taskID int64) (int64, error)
}
