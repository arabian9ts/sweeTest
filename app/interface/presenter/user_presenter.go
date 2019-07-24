package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type UserPresenter struct {
}

func NewUserPresenter() (port.UserOutput) {
	return &UserPresenter{}
}

func (*UserPresenter) HandleGetStudent(student *model.Student, err error) (*dto.GetStudentByIdOutputForm, error) {
	output := &dto.GetStudentByIdOutputForm{
		StudentNo: student.StudentNo,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetTa(ta *model.Ta, err error) (*dto.GetTaByIdOutputForm, error) {
	output := &dto.GetTaByIdOutputForm{
		StudentNo: ta.StudentNo,
		FirstName: ta.FirstName,
		LastName:  ta.LastName,
		Email:     ta.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetTeacher(teacher *model.Teacher, err error) (*dto.GetTeacherByIdOutputForm, error) {
	output := &dto.GetTeacherByIdOutputForm{
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
	}
	return output, err
}

func (*UserPresenter) HandleCreateStudent(id int64, err error) (*dto.CreateStudentOutputForm, error) {
	output := &dto.CreateStudentOutputForm{LastInsertedId: id}
	return output, err
}

func (*UserPresenter) HandleCreateTa(id int64, err error) (*dto.CreateTaOutputForm, error) {
	output := &dto.CreateTaOutputForm{LastInsertedId: id}
	return output, err
}

func (*UserPresenter) HandleCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error) {
	output := &dto.CreateTeacherOutputForm{LastInsertedId: id}
	return output, err
}
