package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type UserPresenter struct {
}

func NewUserPresenter() port.UserOutput {
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

func (*UserPresenter) HandleGetAssistant(assistant *model.Assistant, err error) (*dto.GetAssistantByIdOutputForm, error) {
	output := &dto.GetAssistantByIdOutputForm{
		StudentNo: assistant.StudentNo,
		FirstName: assistant.FirstName,
		LastName:  assistant.LastName,
		Email:     assistant.Email,
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

func (*UserPresenter) HandleCreateAssistant(id int64, err error) (*dto.CreateAssistantOutputForm, error) {
	output := &dto.CreateAssistantOutputForm{LastInsertedId: id}
	return output, err
}

func (*UserPresenter) HandleCreateTeacher(id int64, err error) (*dto.CreateTeacherOutputForm, error) {
	output := &dto.CreateTeacherOutputForm{LastInsertedId: id}
	return output, err
}

func (*UserPresenter) HandleUpdateStudent(count int64, err error) (*dto.UpdateStudentOutputForm, error) {
	updated := count != 0
	return &dto.UpdateStudentOutputForm{
		Updated: updated,
	}, err
}

func (*UserPresenter) HandleUpdateAssistant(count int64, err error) (*dto.UpdateAssistantOutputForm, error) {
	updated := count != 0
	return &dto.UpdateAssistantOutputForm{
		Updated: updated,
	}, err
}

func (*UserPresenter) HandleUpdateTeacher(count int64, err error) (*dto.UpdateTeacherOutputForm, error) {
	updated := count != 0
	return &dto.UpdateTeacherOutputForm{
		Updated: updated,
	}, err
}
