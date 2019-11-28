package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type UserUseCase interface {
	GetUsersUseCase
	GetUserUseCase
	GetParticipatedUsersCase
	CreateUserUseCase
	UpdateUserUseCase
}

type GetUsersUseCase interface {
	GetStudents(limit int, offset int) (dto.GetStudentsOutputForm, error)
	GetAssistants(limit int, offset int) (dto.GetAssistantsOutputForm, error)
	GetTeachers(limit int, offset int) (dto.GetTeachersOutputForm, error)
}

type GetParticipatedUsersCase interface {
	GetParticipatedStudents(lectureID int64, limit int, offset int) (dto.GetStudentsOutputForm, error)
	GetParticipatedAssistants(lectureID int64, limit int, offset int) (dto.GetAssistantsOutputForm, error)
	GetParticipatedTeachers(lectureID int64, limit int, offset int) (dto.GetTeachersOutputForm, error)
}

type GetUserUseCase interface {
	GetStudentById(id int64) (*dto.GetStudentOutputForm, error)
	GetAssistantById(id int64) (*dto.GetAssistantOutputForm, error)
	GetTeacherById(id int64) (*dto.GetTeacherOutputForm, error)
}

type CreateUserUseCase interface {
	CreateStudent(form *dto.CreateStudentInputForm) (*dto.CreateStudentOutputForm, error)
	CreateAssistant(form *dto.CreateAssistantInputForm) (*dto.CreateAssistantOutputForm, error)
	CreateTeacher(form *dto.CreateTeacherInputForm) (*dto.CreateTeacherOutputForm, error)
}

type UpdateUserUseCase interface {
	UpdateStudent(form *dto.UpdateStudentInputForm) (*dto.UpdateStudentOutputForm, error)
	UpdateAssistant(form *dto.UpdateAssistantInputForm) (*dto.UpdateAssistantOutputForm, error)
	UpdateTeacher(form *dto.UpdateTeacherInputForm) (*dto.UpdateTeacherOutputForm, error)
}
