package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserOutput interface {
	GetUserOutput
	CreateUserOutput
	UpdateUserOutput
}

type GetUserOutput interface {
	HandleGetStudent(student *model.Student, err error) (*dto.GetStudentByIdOutputForm, error)
	HandleGetAssistant(assistant *model.Assistant, err error) (*dto.GetAssistantByIdOutputForm, error)
	HandleGetTeacher(teacher *model.Teacher, err error) (*dto.GetTeacherByIdOutputForm, error)
}

type CreateUserOutput interface {
	HandleCreateStudent(id int64, err error) (*dto.CreateStudentOutputForm, error)
	HandleCreateAssistant(id int64, err error) (*dto.CreateAssistantOutputForm, error)
	HandleCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error)
}

type UpdateUserOutput interface {
	HandleUpdateStudent(id int64, err error) (*dto.UpdateStudentOutputForm, error)
	HandleUpdateAssistant(id int64, err error) (*dto.UpdateAssistantOutputForm, error)
	HandleUpdateTeacher(id int64, err error) (*dto.UpdateTeacherOutputForm, error)
}
