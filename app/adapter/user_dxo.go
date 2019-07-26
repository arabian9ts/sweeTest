package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/util"
)

func ConvertStudentInputFormToUser(form *dto.CreateStudentInputForm) (student *model.Student) {
	student = &model.Student{}
	student.StudentNo = form.StudentNo
	student.FirstName = form.FirstName
	student.LastName = form.LastName
	student.Email = form.Email
	student.Digest = util.EncryptPassword(form.Password)
	return
}

func ConvertAssistantInputFormToAssistant(form *dto.CreateAssistantInputForm) (assistant *model.Assistant) {
	assistant = &model.Assistant{}
	assistant.StudentNo = form.StudentNo
	assistant.FirstName = form.FirstName
	assistant.LastName = form.LastName
	assistant.Email = form.Email
	assistant.Digest = util.EncryptPassword(form.Password)
	return
}

func ConvertTeacherInputFormToTeacher(form *dto.CreateTeacherInputForm) (teacher *model.Teacher) {
	teacher = &model.Teacher{}
	teacher.FirstName = form.FirstName
	teacher.LastName = form.LastName
	teacher.Email = form.Email
	teacher.Digest = util.EncryptPassword(form.Password)
	return
}
