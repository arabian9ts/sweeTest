package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type UserRepository interface {
	GetStudents(limit, offset int) (model.Students, error)
	GetAssistants(limit, offset int) (model.Assistants, error)
	GetTeachers(limit, offset int) (model.Teachers, error)

	GetStudentById(int64) (*model.Student, error)
	GetAssistantById(int64) (*model.Assistant, error)
	GetTeacherById(int64) (*model.Teacher, error)
	GetAdminById(int64) (*model.Admin, error)

	GetStudentByStudentNo(studentNo string) (*model.Student, error)
	GetAssistantByStudentNo(studentNo string) (*model.Assistant, error)
	GetTeacherByEmail(email string) (*model.Teacher, error)
	GetAdminByEmail(email string) (*model.Admin, error)

	GetParticipatingStudentsOfLecture(lectureID int64, limit int, offset int) (model.Students, error)
	GetParticipatingAssistantsOfLecture(lectureID int64, limit int, offset int) (model.Assistants, error)
	GetParticipatingTeachersOfLecture(lectureID int64, limit int, offset int) (model.Teachers, error)

	InsertStudent(student *model.Student) (*model.Student, error)
	InsertAssistant(assistant *model.Assistant) (*model.Assistant, error)
	InsertTeacher(teacher *model.Teacher) (*model.Teacher, error)
	InsertAdmin(admin *model.Admin) (*model.Admin, error)

	UpdateStudent(student *model.Student) (*model.Student, error)
	UpdateAssistant(assistant *model.Assistant) (*model.Assistant, error)
	UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error)
}
