package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type AuthOutput interface {
	HandleAuthorizeStudent(student *model.Student, err error) (*dto.AuthorizeStudentOutputForm, error)
	HandleAuthorizeAssistant(assistant *model.Assistant, err error) (*dto.AuthorizeAssistantOutputForm, error)
	HandleAuthorizeTeacher(teacher *model.Teacher, err error) (*dto.AuthorizeTeacherOutputForm, error)
	HandleAuthorizeAdmin(admin *model.Admin, err error) (*dto.AuthorizeAdminOutputForm, error)
}
