package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	GetUserUseCase
	CreateUserUseCase
}

type GetUserUseCase interface {
	GetStudentById(id int64) (*dto.GetStudentByIdOutputForm, error)
	GetAssistantById(id int64) (*dto.GetAssistantByIdOutputForm, error)
	GetTeacherById(id int64) (*dto.GetTeacherByIdOutputForm, error)
}

type CreateUserUseCase interface {
	CreateStudent(form *dto.CreateStudentInputForm) (*dto.CreateStudentOutputForm, error)
	CreateAssistant(form *dto.CreateAssistantInputForm) (*dto.CreateAssistantOutputForm, error)
	CreateTeacher(form *dto.CreateTeacherInputForm) (*dto.CreateTeacherOutputForm, error)
}
