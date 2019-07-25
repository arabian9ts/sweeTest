package mock

import (
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

func (mock *UserRepositoryMock) InsertStudent(student *model.Student) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.StudentFixture.ID, nil
}

func (mock *UserRepositoryMock) InsertAssistant(assistant *model.Assistant) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.AssistantFixture.ID, nil
}

func (mock *UserRepositoryMock) InsertTeacher(teacher *model.Teacher) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.TeacherFixture.ID, nil
}

func (mock *UserRepositoryMock) InsertAdmin(admin *model.Admin) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.AdminFixture.ID, nil
}
