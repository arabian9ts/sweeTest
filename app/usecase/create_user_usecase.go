package usecase

import "github.com/arabian9ts/sweeTest/app/dto"

type CreateUserUseCase interface {
	CreateStudent(data dto.CreateStudentInputForm) dto.CreateStudentOutputForm
	CreateTa(data dto.CreateTaInputForm) dto.CreateTaOutputForm
	CreateTeacher(data dto.CreateTeacherInputForm) dto.CreateTeacherOutputForm
}
