package interactor

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthInteractor struct {
	UserRepository repository.UserRepository
	AuthOutput     port.AuthOutput
}

func NewAuthInteractor(repository repository.UserRepository, output port.AuthOutput) (port.AuthUseCase, error) {
	return &AuthInteractor{UserRepository: repository, AuthOutput: output}, nil
}

func (interactor *AuthInteractor) AuthorizeStudent(studentNo string, password string) (*dto.AuthorizeStudentOutputForm, error) {
	student, err := interactor.UserRepository.GetStudentByStudentNo(studentNo)
	if err != nil {
		return interactor.AuthOutput.HandleAuthorizeStudent(student, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(student.Digest), []byte(password))
	return interactor.AuthOutput.HandleAuthorizeStudent(student, err)
}

func (interactor *AuthInteractor) AuthorizeAssistant(studentNo string, password string) (*dto.AuthorizeAssistantOutputForm, error) {
	assistant, err := interactor.UserRepository.GetAssistantByStudentNo(studentNo)
	if err != nil {
		return interactor.AuthOutput.HandleAuthorizeAssistant(assistant, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(assistant.Digest), []byte(password))
	return interactor.AuthOutput.HandleAuthorizeAssistant(assistant, err)
}

func (interactor *AuthInteractor) AuthorizeTeacher(email string, password string) (*dto.AuthorizeTeacherOutputForm, error) {
	teacher, err := interactor.UserRepository.GetTeacherByEmail(email)
	if err != nil {
		return interactor.AuthOutput.HandleAuthorizeTeacher(teacher, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(teacher.Digest), []byte(password))
	return interactor.AuthOutput.HandleAuthorizeTeacher(teacher, err)
}

func (interactor *AuthInteractor) AuthorizeAdmin(email string, password string) (*dto.AuthorizeAdminOutputForm, error) {
	admin, err := interactor.UserRepository.GetAdminByEmail(email)
	if err != nil {
		return interactor.AuthOutput.HandleAuthorizeAdmin(admin, err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(admin.Digest), []byte(password))
	return interactor.AuthOutput.HandleAuthorizeAdmin(admin, err)
}
