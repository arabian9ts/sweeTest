package interactor

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/util"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserOutput     port.UserOutput
}

func NewUserInteractor(repository repository.UserRepository, output port.UserOutput) (*UserInteractor, error) {
	return &UserInteractor{UserRepository: repository, UserOutput:output}, nil
}

func (interactor *UserInteractor) CreateStudent(form dto.CreateStudentInputForm) (*dto.CreateStudentOutputForm, error) {
	student := &model.Student{StudentNo: form.StudentNo, FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	student.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertStudent(student)
	return interactor.UserOutput.HandlerCreateStudent(id, err)
}

func (interactor *UserInteractor) CreateTa(form dto.CreateTaInputForm) (*dto.CreateTaOutputForm, error) {
	ta := &model.Ta{StudentNo: form.StudentNo, FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	ta.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertTa(ta)
	return interactor.UserOutput.HandlerCreateTa(id, err)
}

func (interactor *UserInteractor) CreateTeacher(form dto.CreateTeacherInputForm) (*dto.CreateTeacherOutputForm, error) {
	teacher := &model.Teacher{FirstName: form.FirstName, LastName: form.LastName, Email: form.Email}
	teacher.Digest = util.EncryptPassword(form.Password)
	id, err := interactor.UserRepository.InsertTeacher(teacher)
	return interactor.UserOutput.HandlerCreateTeacher(id, err)
}
