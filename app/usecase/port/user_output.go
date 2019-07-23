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
	HandleGetStudent(student *model.Student, err error) (*dto.ReadStudentOutputForm, error)
	HandleGetTa(ta *model.Ta, err error) (*dto.ReadTaOutputForm, error)
	HandleGetTeacher(teacher *model.Teacher, err error) (*dto.ReadTeacherOutputForm, error)
}

type UserWriteOutput interface {
	HandleCreateStudent(id int64, err error) (*dto.WriteStudentOutputForm, error)
	HandleCreateTa(id int64, err error) (*dto.WriteTaOutputForm, error)
	HandleCreateTeacher(id int64, err error) (*dto.WriteTeacherOutputForm, error)
}
