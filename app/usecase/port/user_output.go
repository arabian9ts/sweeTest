package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserOutput interface {
	GetUsersOutput
	GetUserOutput
	CreateUserOutput
	UpdateUserOutput
}

type GetUsersOutput interface {
	HandleGetStudents(students []*model.Student, err error) (dto.GetStudentsOutputForm, error)
	HandleGetAssistants(assistants []*model.Assistant, err error) (dto.GetAssistantsOutputForm, error)
	HandleGetTeachers(teachers []*model.Teacher, err error) (dto.GetTeachersOutputForm, error)
}

type GetUserOutput interface {
	HandleGetStudent(student *model.Student, err error) (*dto.GetStudentOutputForm, error)
	HandleGetAssistant(assistant *model.Assistant, err error) (*dto.GetAssistantOutputForm, error)
	HandleGetTeacher(teacher *model.Teacher, err error) (*dto.GetTeacherOutputForm, error)
}

type CreateUserOutput interface {
	HandleCreateStudent(student *model.Student, err error) (*dto.CreateStudentOutputForm, error)
	HandleCreateAssistant(assistant *model.Assistant, err error) (*dto.CreateAssistantOutputForm, error)
	HandleCreateTeacher(teacher *model.Teacher, err error) (*dto.CreateTeacherOutputForm, error)
}

type UpdateUserOutput interface {
	HandleUpdateStudent(student *model.Student, err error) (*dto.UpdateStudentOutputForm, error)
	HandleUpdateAssistant(assistant *model.Assistant, err error) (*dto.UpdateAssistantOutputForm, error)
	HandleUpdateTeacher(teacher *model.Teacher, err error) (*dto.UpdateTeacherOutputForm, error)
}
