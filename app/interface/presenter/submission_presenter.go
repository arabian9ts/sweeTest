package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type SubmissionPresenter struct{}

func NewSubmissionPresenter() port.SubmissionOutput {
	return &SubmissionPresenter{}
}

func (*SubmissionPresenter) HandleGetSubmissionsByStudentID(submissions model.Submissions, err error) (dto.GetSubmissionsOutputForm, error) {
	output := dto.GetSubmissionsOutputForm{}
	for _, submission := range submissions {
		form := &dto.GetSubmissionOutputForm{
			ID:        submission.ID,
			StudentID: submission.StudentID,
			CreatedAt: submission.CreatedAt,
			UpdatedAt: submission.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*SubmissionPresenter) HandleGetSubmissionsByTaskID(submissions model.Submissions, err error) (dto.GetSubmissionsOutputForm, error) {
	output := dto.GetSubmissionsOutputForm{}
	for _, submission := range submissions {
		form := &dto.GetSubmissionOutputForm{
			ID:        submission.ID,
			StudentID: submission.StudentID,
			CreatedAt: submission.CreatedAt,
			UpdatedAt: submission.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*SubmissionPresenter) HandleGetSubmissionTextByID(submissionText model.SubmissionText, err error) (*dto.GetSubmissionTextOutputForm, error) {
	return &dto.GetSubmissionTextOutputForm{
		ID:           submissionText.ID,
		SubmissionID: submissionText.SubmissionID,
		FileName:     submissionText.FileName,
		Contents:     submissionText.Contents,
		CreatedAt:    submissionText.CreatedAt,
		UpdatedAt:    submissionText.UpdatedAt,
	}, err
}

func (*SubmissionPresenter) HandleGetSubmissionTextsBySubmissionID(submissionTexts model.SubmissionTexts, err error) (dto.GetSubmissionTextsOutputForm, error) {
	output := dto.GetSubmissionTextsOutputForm{}
	for _, submissionText := range submissionTexts {
		form := &dto.GetSubmissionTextOutputForm{
			ID:           submissionText.ID,
			SubmissionID: submissionText.SubmissionID,
			FileName:     submissionText.FileName,
			Contents:     submissionText.Contents,
			CreatedAt:    submissionText.CreatedAt,
			UpdatedAt:    submissionText.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (p *SubmissionPresenter) HandleCreateSubmission(submission *model.Submission, err error) (*dto.CreateSubmissionOutputForm, error) {
	return &dto.CreateSubmissionOutputForm{
		ID:        submission.ID,
		StudentID: submission.StudentID,
		CreatedAt: submission.CreatedAt,
		UpdatedAt: submission.UpdatedAt,
	}, err
}

func (*SubmissionPresenter) HandleCreateSubmissionText(submissionText *model.SubmissionText, err error) (*dto.CreateSubmissionTextOutputForm, error) {
	return &dto.CreateSubmissionTextOutputForm{
		ID:           submissionText.ID,
		SubmissionID: submissionText.SubmissionID,
		FileName:     submissionText.FileName,
		Contents:     submissionText.Contents,
		CreatedAt:    submissionText.CreatedAt,
		UpdatedAt:    submissionText.UpdatedAt,
	}, err
}

func (*SubmissionPresenter) HandleUpdateSubmissionText(submissionText *model.SubmissionText, err error) (*dto.UpdateSubmissionTextOutputForm, error) {
	return &dto.UpdateSubmissionTextOutputForm{
		ID:           submissionText.ID,
		SubmissionID: submissionText.SubmissionID,
		FileName:     submissionText.FileName,
		Contents:     submissionText.Contents,
		CreatedAt:    submissionText.CreatedAt,
		UpdatedAt:    submissionText.UpdatedAt,
	}, err
}

func (*SubmissionPresenter) HandleDeleteSubmissionText(id int64, err error) (*dto.DeleteSubmissionTextOutputForm, error) {
	return &dto.DeleteSubmissionTextOutputForm{
		AffectedRowsCount: id,
	}, err
}
