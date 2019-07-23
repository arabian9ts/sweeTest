package presenter

import (
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type UserPresenter struct {
}

func NewUserPresenter() (port.UserOutput) {
	return &UserPresenter{}
}

func (*UserPresenter) HandlerCreateStudent(id int64, err error) (*dto.CreateStudentOutputForm, error) {
	output := &dto.CreateStudentOutputForm{LastCreatedUserId: id}
	return output, err
}

func (*UserPresenter) HandlerCreateTa(id int64, err error) (*dto.CreateTaOutputForm, error) {
	output := &dto.CreateTaOutputForm{LastCreatedUserId: id}
	return output, err
}

func (*UserPresenter) HandlerCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error) {
	output := &dto.CreateTeacherOutputForm{LastCreatedUserId: id}
	return output, err
}
