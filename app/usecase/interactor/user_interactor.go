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

func (interactor *UserInteractor) GetStudentById(id int64) (*dto.ReadStudentOutputForm, error) {
	return interactor.UserOutput.HandleGetStudent(
		interactor.UserRepository.GetStudentById(id),
	)
}

func (interactor *UserInteractor) GetTaById(id int64) (*dto.ReadTaOutputForm, error) {
	return interactor.UserOutput.HandleGetTa(
		interactor.UserRepository.GetTaById(id),
	)
}

func (interactor *UserInteractor) GetTeacherById(id int64) (*dto.ReadTeacherOutputForm, error) {
	return interactor.UserOutput.HandleGetTeacher(
		interactor.UserRepository.GetTeacherById(id),
	)
}

func (interactor *UserInteractor) CreateStudent(student *model.Student) (*dto.WriteStudentOutputForm, error) {
	return interactor.UserOutput.HandleCreateStudent(
		interactor.UserRepository.InsertStudent(student),
	)
}

func (interactor *UserInteractor) CreateTa(ta *model.Ta) (*dto.WriteTaOutputForm, error) {
	return interactor.UserOutput.HandleCreateTa(
		interactor.UserRepository.InsertTa(ta),
	)
}

func (interactor *UserInteractor) CreateTeacher(teacher *model.Teacher) (*dto.WriteTeacherOutputForm, error) {
	return interactor.UserOutput.HandleCreateTeacher(
		interactor.UserRepository.InsertTeacher(teacher),
	)
}
