package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	UserReadUseCase
	UserWriteUseCase
}

type UserReadUseCase interface {
	GetStudentById(id int64) (*dto.ReadStudentOutputForm, error)
	GetTaById(id int64) (*dto.ReadTaOutputForm, error)
	GetTeacherById(id int64) (*dto.ReadTeacherOutputForm, error)
}

type UserWriteUseCase interface {
	CreateStudent(student *model.Student) (*dto.WriteStudentOutputForm, error)
	CreateTa(ta *model.Ta) (*dto.WriteTaOutputForm, error)
	CreateTeacher(teacher *model.Teacher) (*dto.WriteTeacherOutputForm, error)
}
