package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type UserRepository interface {
	GetStudentById(int64) (*model.Student, error)
	GetTaById(int64) (*model.Ta, error)
	GetTeacherById(int64) (*model.Teacher, error)
	GetAdminById(int64) (*model.Admin, error)

	InsertStudent(student *model.Student) (int64, error)
	InsertTa(ta *model.Ta) (int64, error)
	InsertTeacher(teacher *model.Teacher) (int64, error)
	InsertAdmin(admin *model.Admin) (int64, error)
}
