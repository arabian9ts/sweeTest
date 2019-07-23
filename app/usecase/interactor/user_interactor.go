package interactor

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserOutput     port.UserOutput
}

func NewUserInteractor(repository repository.UserRepository, output port.UserOutput) (*UserInteractor, error) {
	return &UserInteractor{UserRepository: repository, UserOutput: output}, nil
}

func (interactor *UserInteractor) CreateStudent(student *model.Student) (*dto.CreateStudentOutputForm, error) {
	return interactor.UserOutput.HandlerCreateStudent(
		interactor.UserRepository.InsertStudent(student),
	)
}

func (interactor *UserInteractor) CreateTa(ta *model.Ta) (*dto.CreateTaOutputForm, error) {
	return interactor.UserOutput.HandlerCreateTa(
		interactor.UserRepository.InsertTa(ta),
	)
}

func (interactor *UserInteractor) CreateTeacher(teacher *model.Teacher) (*dto.CreateTeacherOutputForm, error) {
	return interactor.UserOutput.HandlerCreateTeacher(
		interactor.UserRepository.InsertTeacher(teacher),
	)
}
