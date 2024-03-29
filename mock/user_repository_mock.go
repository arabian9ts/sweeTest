package mock

import (
	"fmt"

	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/fixture"
)

type UserRepositoryMock struct {
	StudentFixture   *model.Student
	AssistantFixture *model.Assistant
	TeacherFixture   *model.Teacher
	AdminFixture     *model.Admin
	Error            error
}

func NewUserRepositoryMock() repository.UserRepository {
	return &UserRepositoryMock{
		StudentFixture:   fixture.NewValidStudent(),
		AssistantFixture: fixture.NewValidAssistant(),
		TeacherFixture:   fixture.NewValidTeacher(),
		AdminFixture:     fixture.NewValidAdmin(),
		Error:            nil,
	}
}

func (mock *UserRepositoryMock) GetStudentById(int64) (*model.Student, error) {
	if mock.Error != nil {
		return &model.Student{}, mock.Error
	}
	return mock.StudentFixture, nil
}

func (mock *UserRepositoryMock) GetAssistantById(int64) (*model.Assistant, error) {
	if mock.Error != nil {
		return &model.Assistant{}, mock.Error
	}
	return mock.AssistantFixture, nil
}

func (mock *UserRepositoryMock) GetTeacherById(int64) (*model.Teacher, error) {
	if mock.Error != nil {
		return &model.Teacher{}, mock.Error
	}
	return mock.TeacherFixture, nil
}

func (mock *UserRepositoryMock) GetAdminById(int64) (*model.Admin, error) {
	if mock.Error != nil {
		return &model.Admin{}, mock.Error
	}
	return mock.AdminFixture, nil
}

func (mock *UserRepositoryMock) GetStudents(limit, offset int) (model.Students, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Student{mock.StudentFixture}, nil
}

func (mock *UserRepositoryMock) GetAssistants(limit, offset int) (model.Assistants, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Assistant{mock.AssistantFixture}, nil
}

func (mock *UserRepositoryMock) GetTeachers(limit, offset int) (model.Teachers, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Teacher{mock.TeacherFixture}, nil
}

func (mock *UserRepositoryMock) GetParticipatingStudentsOfLecture(lectureID int64, limit int, offset int) (model.Students, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Student{mock.StudentFixture}, nil
}

func (mock *UserRepositoryMock) GetParticipatingAssistantsOfLecture(lectureID int64, limit int, offset int) (model.Assistants, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Assistant{mock.AssistantFixture}, nil
}

func (mock *UserRepositoryMock) GetParticipatingTeachersOfLecture(lectureID int64, limit int, offset int) (model.Teachers, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Teacher{mock.TeacherFixture}, nil
}

func (mock *UserRepositoryMock) InsertStudent(student *model.Student) (*model.Student, error) {
	if mock.Error != nil {
		return &model.Student{}, mock.Error
	}
	return mock.StudentFixture, nil
}

func (mock *UserRepositoryMock) InsertAssistant(assistant *model.Assistant) (*model.Assistant, error) {
	if mock.Error != nil {
		return &model.Assistant{}, mock.Error
	}
	return mock.AssistantFixture, nil
}

func (mock *UserRepositoryMock) InsertTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	if mock.Error != nil {
		return &model.Teacher{}, mock.Error
	}
	return mock.TeacherFixture, nil
}

func (mock *UserRepositoryMock) InsertAdmin(admin *model.Admin) (*model.Admin, error) {
	if mock.Error != nil {
		return admin, mock.Error
	}
	return mock.AdminFixture, nil
}

func (mock *UserRepositoryMock) GetStudentByStudentNo(studentNo string) (*model.Student, error) {
	if mock.Error != nil {
		return &model.Student{}, mock.Error
	}
	if mock.StudentFixture.StudentNo != studentNo {
		return &model.Student{}, mock.Error
	}
	return mock.StudentFixture, nil
}

func (mock *UserRepositoryMock) GetAssistantByStudentNo(studentNo string) (*model.Assistant, error) {
	if mock.Error != nil {
		return &model.Assistant{}, mock.Error
	}
	if mock.AssistantFixture.StudentNo != studentNo {
		return &model.Assistant{}, mock.Error
	}
	return mock.AssistantFixture, nil
}

func (mock *UserRepositoryMock) GetTeacherByEmail(email string) (*model.Teacher, error) {
	if mock.Error != nil {
		return &model.Teacher{}, mock.Error
	}
	if mock.TeacherFixture.Email != email {
		return &model.Teacher{}, mock.Error
	}
	return mock.TeacherFixture, nil
}

func (mock *UserRepositoryMock) GetAdminByEmail(email string) (*model.Admin, error) {
	if mock.Error != nil {
		return &model.Admin{}, mock.Error
	}
	if mock.AdminFixture.Email != email {
		return &model.Admin{}, mock.Error
	}
	return mock.AdminFixture, nil
}

func (mock *UserRepositoryMock) UpdateStudent(student *model.Student) (*model.Student, error) {
	if mock.Error != nil {
		return &model.Student{}, mock.Error
	}
	if student.ID != mock.StudentFixture.ID {
		return &model.Student{}, fmt.Errorf("student not exists")
	}
	return student, nil
}

func (mock *UserRepositoryMock) UpdateAssistant(assistant *model.Assistant) (*model.Assistant, error) {
	if mock.Error != nil {
		return &model.Assistant{}, mock.Error
	}
	if assistant.ID != mock.AssistantFixture.ID {
		return &model.Assistant{}, fmt.Errorf("assistant not exists")
	}
	return assistant, nil
}

func (mock *UserRepositoryMock) UpdateTeacher(teacher *model.Teacher) (*model.Teacher, error) {
	if mock.Error != nil {
		return &model.Teacher{}, mock.Error
	}
	if teacher.ID != mock.TeacherFixture.ID {
		return &model.Teacher{}, fmt.Errorf("teacher not exists")
	}
	return teacher, nil
}
