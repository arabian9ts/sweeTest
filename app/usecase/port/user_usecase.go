package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	GetUserUseCase
	CreateUserUseCase
}

type GetUserUseCase interface {
	GetStudentById(id int64) (*dto.GetStudentByIdOutputForm, error)
	GetTaById(id int64) (*dto.GetTaByIdOutputForm, error)
	GetTeacherById(id int64) (*dto.GetTeacherByIdOutputForm, error)
}

type CreateUserUseCase interface {
	CreateStudent(student *model.Student) (*dto.CreateStudentOutputForm, error)
	CreateTa(ta *model.Ta) (*dto.CreateTaOutputForm, error)
	CreateTeacher(teacher *model.Teacher) (*dto.CreateTeacherOutputForm, error)
}
