package port

import "github.com/arabian9ts/sweeTest/app/dto"

type AuthUseCase interface {
	AuthorizeStudent(studentNo string, password string) (*dto.AuthorizeStudentOutputForm, error)
	AuthorizeAssistant(studentNo string, password string) (*dto.AuthorizeAssistantOutputForm, error)
	AuthorizeTeacher(email string, password string) (*dto.AuthorizeTeacherOutputForm, error)
	AuthorizeAdmin(email string, password string) (*dto.AuthorizeAdminOutputForm, error)
}
