package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
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

func (interactor *LectureInteractor) GetParticipatingLecturesOfStudent(studentID int64, limit int, offset int) (dto.GetLecturesOutputForm, error) {
	return interactor.LectureOutput.HandleGetLectures(
		interactor.LectureRepository.GetParticipatedLecturesOfStudent(studentID, limit, offset),
	)
}

func (interactor *LectureInteractor) GetParticipatingLecturesOfAssistant(studentID int64, limit int, offset int) (dto.GetLecturesOutputForm, error) {
	return interactor.LectureOutput.HandleGetLectures(
		interactor.LectureRepository.GetParticipatedLecturesOfAssistant(studentID, limit, offset),
	)
}

func (interactor *LectureInteractor) GetParticipatingLecturesOfTeacher(teacherID int64, limit int, offset int) (dto.GetLecturesOutputForm, error) {
	return interactor.LectureOutput.HandleGetLectures(
		interactor.LectureRepository.GetParticipatedLecturesOfTeacher(teacherID, limit, offset),
	)
}

func (interactor *LectureInteractor) CreateLecture(form *dto.CreateLectureInputForm) (*dto.CreateLectureOutputForm, error) {
	lecture := adapter.ConvertCreateLectureInputFormToLecture(form)
	return interactor.LectureOutput.HandleCreateLecture(
		interactor.LectureRepository.CreateLecture(lecture),
	)
}

func (interactor *LectureInteractor) UpdateLecture(form *dto.UpdateLectureInputForm) (*dto.UpdateLectureOutputForm, error) {
	lecture := adapter.ConvertUpdateLectureInputFormToLecture(form)
	return interactor.LectureOutput.HandleUpdateLecture(
		interactor.LectureRepository.UpdateLecture(lecture),
	)
}

func (interactor *LectureInteractor) DeleteLecture(id int64) (*dto.DeleteLectureOutputForm, error) {
	return interactor.LectureOutput.HandleDeleteLecture(
		interactor.LectureRepository.DeleteLecture(id),
	)
}
