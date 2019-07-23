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
	student.Digest = util.EncryptPassword(form.Password) // error handling
	return
}

func ConvertTaInputFormToTa(form *dto.CreateTaInputForm) (ta *model.Ta) {
	ta = &model.Ta{}
	ta.StudentNo = form.StudentNo
	ta.FirstName = form.FirstName
	ta.LastName = form.LastName
	ta.Email = form.Email
	ta.Digest = util.EncryptPassword(form.Password)
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