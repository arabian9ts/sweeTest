package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type ParticipationInteractor struct {
	ParticipationRepository repository.ParticipationRepository
	ParticipationOutput     port.ParticipationOutput
}

func NewParticipationInteractor(repository repository.ParticipationRepository, output port.ParticipationOutput) (*ParticipationInteractor, error) {
	return &ParticipationInteractor{ParticipationRepository: repository, ParticipationOutput: output}, nil
}

func (interactor *ParticipationInteractor) CreateStudentLecture(form *dto.CreateStudentLectureInputForm) (*dto.CreateStudentLectureOutputForm, error) {
	studentLecture := adapter.ConvertCreateStudentLectureInputFormToStudentLecture(form)
	return interactor.ParticipationOutput.HandleCreateStudentLecture(
		interactor.ParticipationRepository.CreateStudentLecture(studentLecture),
	)
}

func (interactor *ParticipationInteractor) CreateAssistantLecture(form *dto.CreateAssistantLectureInputForm) (*dto.CreateAssistantLectureOutputForm, error) {
	assistantLecture := adapter.ConvertCreateAssistantLectureInputFormToAssistantLecture(form)
	return interactor.ParticipationOutput.HandleCreateAssistantLecture(
		interactor.ParticipationRepository.CreateAssistantLecture(assistantLecture),
	)
}

func (interactor *ParticipationInteractor) CreateTeacherLecture(form *dto.CreateTeacherLectureInputForm) (*dto.CreateTeacherLectureOutputForm, error) {
	teacherLecture := adapter.ConvertCreateTeacherLectureInputFormToTeacherLecture(form)
	return interactor.ParticipationOutput.HandleCreateTeacherLecture(
		interactor.ParticipationRepository.CreateTeacherLecture(teacherLecture),
	)
}

func (interactor *ParticipationInteractor) DeleteStudentLecture(studentID int64, lectureID int64) (*dto.DeleteStudentLectureOutputForm, error) {
	return interactor.ParticipationOutput.HandleDeleteStudentLecture(
		interactor.ParticipationRepository.DeleteStudentLecture(studentID, lectureID),
	)
}

func (interactor *ParticipationInteractor) DeleteAssistantLecture(studentID int64, lectureID int64) (*dto.DeleteAssistantLectureOutputForm, error) {
	return interactor.ParticipationOutput.HandleDeleteAssistantLecture(
		interactor.ParticipationRepository.DeleteAssistantLecture(studentID, lectureID),
	)
}

func (interactor *ParticipationInteractor) DeleteTeacherLecture(teacherID int64, lectureID int64) (*dto.DeleteTeacherLectureOutputForm, error) {
	return interactor.ParticipationOutput.HandleDeleteTeacherLecture(
		interactor.ParticipationRepository.DeleteTeacherLecture(teacherID, lectureID),
	)
}
