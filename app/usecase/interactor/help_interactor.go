package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type HelpInteractor struct {
	HelpOutput     port.HelpOutput
	HelpRepository repository.HelpRepository
}

func NewHelpInteractor(output port.HelpOutput, repository repository.HelpRepository) (*HelpInteractor, error) {
	return &HelpInteractor{HelpOutput: output, HelpRepository: repository}, nil
}

func (interactor *HelpInteractor) GetHelpsByLectureID(lectureID int64, limit int, offset int) (dto.GetHelpsOutputForm, error) {
	return interactor.HelpOutput.HandleGetHelpsByLectureID(
		interactor.HelpRepository.GetHelpsByLectureID(lectureID, limit, offset),
	)
}

func (interactor *HelpInteractor) CreateHelp(form *dto.CreateHelpInputForm) (*dto.CreateHelpOutputForm, error) {
	help := adapter.ConvertCreateHelpInputFormToHelp(form)
	return interactor.HelpOutput.HandleCreateHelp(
		interactor.HelpRepository.CreateHelp(help),
	)
}

func (interactor *HelpInteractor) UpdateHelp(form *dto.UpdateHelpInputForm) (*dto.UpdateHelpOutputForm, error) {
	help := adapter.ConvertUpdateHelpInputFormToHelp(form)
	return interactor.HelpOutput.HandleUpdateHelp(
		interactor.HelpRepository.UpdateHelp(help),
	)
}

func (interactor *HelpInteractor) DeleteHelp(id int64, lectureID int64, studentID int64) (*dto.DeleteHelpOutputForm, error) {
	return interactor.HelpOutput.HandleDeleteHelp(
		interactor.HelpRepository.DeleteHelp(id, lectureID, studentID),
	)
}
