package interactor

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/repository"
	"github.com/arabian9ts/sweeTest/app/util"
)

type CreateUserInteractor struct {
	UserRepository repository.UserRepository
}

func NewCreateUserInteractor(repository repository.UserRepository) (*CreateUserInteractor, error) {
	return &CreateUserInteractor{UserRepository: repository}, nil
}

func (interactor *CreateUserInteractor) CreateStudent(form dto.CreateStudentInputForm) (output dto.CreateStudentOutputForm) {
	student := &model.Student{StudentNo: form.StudentNo, FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	student.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertStudent(student)
	if err != nil {
		return dto.CreateStudentOutputForm{}
	}
	output.LastCreatedUserId = id
	// ToDo: jwt token
	output.JwtToken = "jwt"

	return
}

func (interactor *CreateUserInteractor) CreateTa(form dto.CreateTaInputForm) (output dto.CreateTaOutputForm) {
	ta := &model.Student{StudentNo: form.StudentNo, FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	ta.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertStudent(ta)
	if err != nil {
		return dto.CreateTaOutputForm{}
	}
	output.LastCreatedUserId = id
	// ToDo: jwt token
	output.JwtToken = "jwt"

	return
}

func (interactor *CreateUserInteractor) CreateTeacher(form dto.CreateTeacherInputForm) (output dto.CreateTeacherOutputForm) {
	teacher := &model.Student{FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	teacher.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertStudent(teacher)
	if err != nil {
		return dto.CreateTeacherOutputForm{}
	}
	output.LastCreatedUserId = id
	// ToDo: jwt token
	output.JwtToken = "jwt"

	return
}
