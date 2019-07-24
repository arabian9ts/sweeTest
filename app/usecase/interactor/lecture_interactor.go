package interactor

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type LectureInteractor struct {
	LectureRepository repository.LectureRepository
	LectureOutput     port.LectureOutput
}

func NewLectureInteractor(repository repository.LectureRepository, output port.LectureOutput) (*LectureInteractor, error) {
	return &LectureInteractor{LectureRepository: repository, LectureOutput: output}, nil
}

func (interactor *LectureInteractor) GetLectureById(id int64) (*dto.GetLectureByIdOutputForm, error) {
	return interactor.LectureOutput.HandleGetLectureById(
		interactor.LectureRepository.GetLectureById(id),
	)
}

func (interactor *LectureInteractor) GetLectures(limit int, offset int) (dto.GetLecturesOutputForm, error) {
	return interactor.LectureOutput.HandleGetLectures(
		interactor.LectureRepository.GetLectures(limit, offset),
	)
}

func (interactor *LectureInteractor) CreateLecture(lecture *model.Lecture) (*dto.CreateLectureOutputForm, error) {
	return interactor.LectureOutput.HandleCreateLecture(
		interactor.LectureRepository.CreateLecture(lecture),
	)
}

func (interactor *LectureInteractor) UpdateLecture(lecture *model.Lecture) (*dto.UpdateLectureOutputForm, error) {
	return interactor.LectureOutput.HandleUpdateLecture(
		interactor.LectureRepository.UpdateLecture(lecture),
	)
}

func (interactor *LectureInteractor) DeleteLecture(id int64) (*dto.DeleteLectureOutputForm, error) {
	return interactor.LectureOutput.HandleDeleteLecture(
		interactor.LectureRepository.DeleteLecture(id),
	)
}
