package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/util"
)

type AuthPresenter struct{}

func NewAuthPresenter() port.AuthOutput {
	return &AuthPresenter{}
}

func (*AuthPresenter) HandleAuthorizeStudent(student *model.Student, err error) (*dto.AuthorizeStudentOutputForm, error) {
	return &dto.AuthorizeStudentOutputForm{
		JwtToken: util.EncodeStudent(student),
	}, err
}

func (*AuthPresenter) HandleAuthorizeAssistant(assistant *model.Assistant, err error) (*dto.AuthorizeAssistantOutputForm, error) {
	return &dto.AuthorizeAssistantOutputForm{
		JwtToken: util.EncodeAssistant(assistant),
	}, err
}

func (*AuthPresenter) HandleAuthorizeTeacher(teacher *model.Teacher, err error) (*dto.AuthorizeTeacherOutputForm, error) {
	return &dto.AuthorizeTeacherOutputForm{
		JwtToken: util.EncodeTeacher(teacher),
	}, err
}

func (*AuthPresenter) HandleAuthorizeAdmin(admin *model.Admin, err error) (*dto.AuthorizeAdminOutputForm, error) {
	return &dto.AuthorizeAdminOutputForm{
		JwtToken: util.EncodeAdmin(admin),
	}, err
}
