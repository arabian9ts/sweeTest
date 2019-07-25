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

func (interactor *UserInteractor) GetStudentById(id int64) (*dto.GetStudentByIdOutputForm, error) {
	return interactor.UserOutput.HandleGetStudent(
		interactor.UserRepository.GetStudentById(id),
	)
}

func (interactor *UserInteractor) GetAssistantById(id int64) (*dto.GetAssistantByIdOutputForm, error) {
	return interactor.UserOutput.HandleGetAssistant(
		interactor.UserRepository.GetAssistantById(id),
	)
}

func (interactor *UserInteractor) GetTeacherById(id int64) (*dto.GetTeacherByIdOutputForm, error) {
	return interactor.UserOutput.HandleGetTeacher(
		interactor.UserRepository.GetTeacherById(id),
	)
}

func (interactor *UserInteractor) CreateStudent(student *model.Student) (*dto.CreateStudentOutputForm, error) {
	return interactor.UserOutput.HandleCreateStudent(
		interactor.UserRepository.InsertStudent(student),
	)
}

func (interactor *UserInteractor) CreateAssistant(assistant *model.Assistant) (*dto.CreateAssistantOutputForm, error) {
	return interactor.UserOutput.HandleCreateAssistant(
		interactor.UserRepository.InsertAssistant(assistant),
	)
}

func (interactor *UserInteractor) CreateTeacher(teacher *model.Teacher) (*dto.CreateTeacherOutputForm, error) {
	return interactor.UserOutput.HandleCreateTeacher(
		interactor.UserRepository.InsertTeacher(teacher),
	)
}
