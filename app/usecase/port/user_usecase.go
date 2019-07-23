package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	CreateStudent(student *model.Student) (*dto.CreateStudentOutputForm, error)
	CreateTa(ta *model.Ta) (*dto.CreateTaOutputForm, error)
	CreateTeacher(teacher *model.Teacher) (*dto.CreateTeacherOutputForm, error)
}
