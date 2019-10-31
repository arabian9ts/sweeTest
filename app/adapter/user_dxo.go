package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/util"
)

func ConvertCreateStudentInputFormToUser(form *dto.CreateStudentInputForm) (student *model.Student) {
	student = &model.Student{}
	student.StudentNo = form.StudentNo
	student.FirstName = form.FirstName
	student.LastName = form.LastName
	student.Email = form.Email
	student.Digest = util.EncryptPassword(form.Password)
	return
}

func ConvertCreateAssistantInputFormToAssistant(form *dto.CreateAssistantInputForm) (assistant *model.Assistant) {
	assistant = &model.Assistant{}
	assistant.StudentNo = form.StudentNo
	assistant.FirstName = form.FirstName
	assistant.LastName = form.LastName
	assistant.Email = form.Email
	assistant.Digest = util.EncryptPassword(form.Password)
	return
}

func ConvertCreateTeacherInputFormToTeacher(form *dto.CreateTeacherInputForm) (teacher *model.Teacher) {
	teacher = &model.Teacher{}
	teacher.FirstName = form.FirstName
	teacher.LastName = form.LastName
	teacher.Email = form.Email
	teacher.Digest = util.EncryptPassword(form.Password)
	return
}

func ConvertUpdateStudentInputFormToUser(form *dto.UpdateStudentInputForm) (student *model.Student) {
	student = &model.Student{}
	student.ID = form.ID
	student.StudentNo = form.StudentNo
	student.FirstName = form.FirstName
	student.LastName = form.LastName
	student.Email = form.Email
	return
}

func ConvertUpdateAssistantInputFormToAssistant(form *dto.UpdateAssistantInputForm) (assistant *model.Assistant) {
	assistant = &model.Assistant{}
	assistant.ID = form.ID
	assistant.StudentNo = form.StudentNo
	assistant.FirstName = form.FirstName
	assistant.LastName = form.LastName
	assistant.Email = form.Email
	return
}

func ConvertUpdateTeacherInputFormToTeacher(form *dto.UpdateTeacherInputForm) (teacher *model.Teacher) {
	teacher = &model.Teacher{}
	teacher.ID = form.ID
	teacher.FirstName = form.FirstName
	teacher.LastName = form.LastName
	teacher.Email = form.Email
	return
}
