package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type SubmissionInteractor struct {
	SubmissionRepository repository.SubmissionRepository
	SubmissionOutput     port.SubmissionOutput
}

func NewSubmissionInteractor(repository repository.SubmissionRepository, output port.SubmissionOutput) (*SubmissionInteractor, error) {
	return &SubmissionInteractor{SubmissionRepository: repository, SubmissionOutput: output}, nil
}

func (interactor *SubmissionInteractor) GetSubmissionTextById(submissionTextID int64) (*dto.GetSubmissionTextOutputForm, error) {
	return interactor.SubmissionOutput.HandleGetSubmissionTextByID(
		interactor.SubmissionRepository.GetSubmissionTextByID(submissionTextID),
	)
}

func (interactor *SubmissionInteractor) GetSubmissionTextsBySubmissionId(submissionID int64, limit int, offset int) (dto.GetSubmissionTextsOutputForm, error) {
	return interactor.SubmissionOutput.HandleGetSubmissionTextsBySubmissionID(
		interactor.SubmissionRepository.GetSubmissionTextsBySubmissionID(submissionID, limit, offset),
	)
}
func (interactor *SubmissionInteractor) CreateSubmission(form *dto.CreateSubmissionInputForm) (*dto.CreateSubmissionOutputForm, error) {
	submission := adapter.ConvertCreateSubmissionInputFormToSubmission(form)
	return interactor.SubmissionOutput.HandleCreateSubmission(
		interactor.SubmissionRepository.CreateSubmission(submission),
	)
}

func (interactor *SubmissionInteractor) CreateSubmissionText(form *dto.CreateSubmissionTextInputForm) (*dto.CreateSubmissionTextOutputForm, error) {
	submissionText := adapter.ConvertCreateSubmissionTextInputFormToSubmissionText(form)
	return interactor.SubmissionOutput.HandleCreateSubmissionText(
		interactor.SubmissionRepository.CreateSubmissionText(submissionText),
	)
}

func (interactor *SubmissionInteractor) UpdateSubmissionText(form *dto.UpdateSubmissionTextInputForm) (*dto.UpdateSubmissionTextOutputForm, error) {
	submissionText := adapter.ConvertUpdateSubmissionTextInputFormToSubmissionText(form)
	return interactor.SubmissionOutput.HandleUpdateSubmissionText(
		interactor.SubmissionRepository.UpdateSubmissionText(submissionText),
	)
}

func (interactor *SubmissionInteractor) DeleteSubmission(id int64, submissionID int64) (*dto.DeleteSubmissionTextOutputForm, error) {
	return interactor.SubmissionOutput.HandleDeleteSubmissionText(
		interactor.SubmissionRepository.DeleteSubmissionText(id, submissionID),
	)
}
