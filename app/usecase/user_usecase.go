package usecase

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	CreateStudent(data dto.CreateStudentInputForm) (*dto.CreateStudentOutputForm, error)
	CreateTa(data dto.CreateTaInputForm) (*dto.CreateTaOutputForm, error)
	CreateTeacher(data dto.CreateTeacherInputForm) (*dto.CreateTeacherOutputForm, error)
}
