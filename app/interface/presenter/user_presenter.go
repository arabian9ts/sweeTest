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

func (*UserPresenter) HandleGetStudents(students []*model.Student, err error) (dto.GetStudentsOutputForm, error) {
	output := dto.GetStudentsOutputForm{}
	for _, student := range students {
		form := &dto.GetStudentOutputForm{
			ID:        student.ID,
			StudentNo: student.StudentNo,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			CreatedAt: student.CreatedAt,
			UpdatedAt: student.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*UserPresenter) HandleGetAssistants(assistants []*model.Assistant, err error) (dto.GetAssistantsOutputForm, error) {
	output := dto.GetAssistantsOutputForm{}
	for _, assistant := range assistants {
		form := &dto.GetAssistantOutputForm{
			ID:        assistant.ID,
			StudentNo: assistant.StudentNo,
			FirstName: assistant.FirstName,
			LastName:  assistant.LastName,
			CreatedAt: assistant.CreatedAt,
			UpdatedAt: assistant.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*UserPresenter) HandleGetTeachers(teachers []*model.Teacher, err error) (dto.GetTeachersOutputForm, error) {
	output := dto.GetTeachersOutputForm{}
	for _, teacher := range teachers {
		form := &dto.GetTeacherOutputForm{
			ID:        teacher.ID,
			FirstName: teacher.FirstName,
			LastName:  teacher.LastName,
			CreatedAt: teacher.CreatedAt,
			UpdatedAt: teacher.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*UserPresenter) HandleGetStudent(student *model.Student, err error) (*dto.GetStudentOutputForm, error) {
	output := &dto.GetStudentOutputForm{
		StudentNo: student.StudentNo,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetAssistant(assistant *model.Assistant, err error) (*dto.GetAssistantOutputForm, error) {
	output := &dto.GetAssistantOutputForm{
		StudentNo: assistant.StudentNo,
		FirstName: assistant.FirstName,
		LastName:  assistant.LastName,
		Email:     assistant.Email,
	}
	return output, err
}

func (*UserPresenter) HandleGetTeacher(teacher *model.Teacher, err error) (*dto.GetTeacherOutputForm, error) {
	output := &dto.GetTeacherOutputForm{
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
	}
	return output, err
}

func (*UserPresenter) HandleCreateStudent(student *model.Student, err error) (*dto.CreateStudentOutputForm, error) {
	output := &dto.CreateStudentOutputForm{
		ID:        student.ID,
		StudentNo: student.StudentNo,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}
	return output, err
}

func (*UserPresenter) HandleCreateAssistant(assistant *model.Assistant, err error) (*dto.CreateAssistantOutputForm, error) {
	output := &dto.CreateAssistantOutputForm{
		ID:        assistant.ID,
		StudentNo: assistant.StudentNo,
		FirstName: assistant.FirstName,
		LastName:  assistant.LastName,
		Email:     assistant.Email,
		CreatedAt: assistant.CreatedAt,
		UpdatedAt: assistant.UpdatedAt,
	}
	return output, err
}

func (*UserPresenter) HandleCreateTeacher(teacher *model.Teacher, err error) (*dto.CreateTeacherOutputForm, error) {
	output := &dto.CreateTeacherOutputForm{
		ID:        teacher.ID,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
	}
	return output, err
}

func (*UserPresenter) HandleUpdateStudent(student *model.Student, err error) (*dto.UpdateStudentOutputForm, error) {
	output := &dto.UpdateStudentOutputForm{
		ID:        student.ID,
		StudentNo: student.StudentNo,
		FirstName: student.FirstName,
		LastName:  student.LastName,
		Email:     student.Email,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
	}
	return output, err
}

func (*UserPresenter) HandleUpdateAssistant(assistant *model.Assistant, err error) (*dto.UpdateAssistantOutputForm, error) {
	output := &dto.UpdateAssistantOutputForm{
		ID:        assistant.ID,
		StudentNo: assistant.StudentNo,
		FirstName: assistant.FirstName,
		LastName:  assistant.LastName,
		Email:     assistant.Email,
		CreatedAt: assistant.CreatedAt,
		UpdatedAt: assistant.UpdatedAt,
	}
	return output, err
}

func (*UserPresenter) HandleUpdateTeacher(teacher *model.Teacher, err error) (*dto.UpdateTeacherOutputForm, error) {
	output := &dto.UpdateTeacherOutputForm{
		ID:        teacher.ID,
		FirstName: teacher.FirstName,
		LastName:  teacher.LastName,
		Email:     teacher.Email,
		CreatedAt: teacher.CreatedAt,
		UpdatedAt: teacher.UpdatedAt,
	}
	return output, err
}
