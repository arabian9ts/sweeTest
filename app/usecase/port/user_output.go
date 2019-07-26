package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserOutput interface {
	UserReadOutput
	UserWriteOutput
}

type UserReadOutput interface {
	HandleGetStudent(student *model.Student, err error) (*dto.GetStudentByIdOutputForm, error)
	HandleGetAssistant(assistant *model.Assistant, err error) (*dto.GetAssistantByIdOutputForm, error)
	HandleGetTeacher(teacher *model.Teacher, err error) (*dto.GetTeacherByIdOutputForm, error)
}

type UserWriteOutput interface {
	HandleCreateStudent(id int64, err error) (*dto.CreateStudentOutputForm, error)
	HandleCreateAssistant(id int64, err error) (*dto.CreateAssistantOutputForm, error)
	HandleCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error)
}
