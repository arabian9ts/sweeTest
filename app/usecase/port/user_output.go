package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserOutput interface {
	HandlerCreateStudent(id int64, err error) (*dto.CreateStudentOutputForm, error)
	HandlerCreateTa(id int64, err error) (*dto.CreateTaOutputForm, error)
	HandlerCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error)
}