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

func (*UserPresenter) HandleGetStudent(student *model.Student, err error) (*dto.ReadStudentOutputForm, error) {
	output := &dto.ReadStudentOutputForm{
		StudentNo: student.StudentNo,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetTa(ta *model.Ta, err error) (*dto.ReadTaOutputForm, error) {
	output := &dto.ReadTaOutputForm{
		StudentNo: ta.StudentNo,
		FirstName: ta.FirstName,
		LastName:  ta.LastName,
		Email:     ta.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetTeacher(teacher *model.Teacher, err error) (*dto.ReadTeacherOutputForm, error) {
	output := &dto.ReadTeacherOutputForm{
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
	}
	return output, err
}

func (*UserPresenter) HandleCreateStudent(id int64, err error) (*dto.WriteStudentOutputForm, error) {
	output := &dto.WriteStudentOutputForm{LastChangedUserId: id}
	return output, err
}

func (*UserPresenter) HandleCreateTa(id int64, err error) (*dto.WriteTaOutputForm, error) {
	output := &dto.WriteTaOutputForm{LastChangedUserId: id}
	return output, err
}

func (*UserPresenter) HandleCreateTeacher(id int64, err error) (*dto.WriteTeacherOutputForm, error) {
	output := &dto.WriteTeacherOutputForm{LastChangedUserId: id}
	return output, err
}
