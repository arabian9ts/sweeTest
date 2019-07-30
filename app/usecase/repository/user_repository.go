package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type UserRepository interface {
	GetStudentById(int64) (*model.Student, error)
	GetAssistantById(int64) (*model.Assistant, error)
	GetTeacherById(int64) (*model.Teacher, error)
	GetAdminById(int64) (*model.Admin, error)

	GetStudentByStudentNo(studentNo string) (*model.Student, error)
	GetAssistantByStudentNo(studentNo string) (*model.Assistant, error)
	GetTeacherByEmail(email string) (*model.Teacher, error)
	GetAdminByEmail(email string) (*model.Admin, error)

	InsertStudent(student *model.Student) (int64, error)
	InsertAssistant(assistant *model.Assistant) (int64, error)
	InsertTeacher(teacher *model.Teacher) (int64, error)
	InsertAdmin(admin *model.Admin) (int64, error)
}
