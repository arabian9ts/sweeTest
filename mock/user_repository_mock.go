package mock

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/fixture"
)

type UserRepositoryMock struct {
	StudentFixture *model.Student
	TaFixture      *model.Ta
	TeacherFixture *model.Teacher
	AdminFixture   *model.Admin
	Error          error
}

func NewUserRepositoryMock() repository.UserRepository {
	return &UserRepositoryMock{
		StudentFixture: fixture.NewValidStudent(),
		TaFixture:      fixture.NewValidTa(),
		TeacherFixture: fixture.NewValidTeacher(),
		AdminFixture:   fixture.NewValidAdmin(),
		Error:          nil,
	}
}

func (mock *UserRepositoryMock) GetStudentById(int64) (*model.Student, error) {
	if mock.Error != nil {
		return &model.Student{}, mock.Error
	}
	return mock.StudentFixture, nil
}

func (mock *UserRepositoryMock) GetTaById(int64) (*model.Ta, error) {
	if mock.Error != nil {
		return &model.Ta{}, mock.Error
	}
	return mock.TaFixture, nil
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

func (mock *UserRepositoryMock) InsertTa(ta *model.Ta) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	return mock.TaFixture.ID, nil
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
