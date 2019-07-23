package interactor

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/interface/presenter"
	"github.com/arabian9ts/sweeTest/fixture"
	"github.com/arabian9ts/sweeTest/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	t.Run("user=student, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		student := fixture.NewValidStudent()
		outputForm, err := userInteractor.CreateStudent(student)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.NotEqual(t, outputForm.LastChangedUserId, 0)
	})

	t.Run("user=student, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		student := fixture.NewValidStudent()
		outputForm, err := userInteractor.CreateStudent(student)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.Equal(t, outputForm.LastChangedUserId, int64(0))
	})
}

func TestCreateTa(t *testing.T) {
	t.Run("user=ta, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		ta := fixture.NewValidTa()
		outputForm, err := userInteractor.CreateTa(ta)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.NotEqual(t, outputForm.LastChangedUserId, 0)
	})

	t.Run("user=ta, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TaFixture: fixture.NewValidTa(),
			Error:     fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		ta := fixture.NewValidTa()
		outputForm, err := userInteractor.CreateTa(ta)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.Equal(t, outputForm.LastChangedUserId, int64(0))
	})
}

func TestCreateTeacher(t *testing.T) {
	t.Run("user=teacher, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		teacher := fixture.NewValidTeacher()
		outputForm, err := userInteractor.CreateTeacher(teacher)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.NotEqual(t, outputForm.LastChangedUserId, 0)
	})

	t.Run("user=teacher, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		teacher := fixture.NewValidTeacher()
		outputForm, err := userInteractor.CreateTeacher(teacher)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastChangedUserId)
		assert.Equal(t, outputForm.LastChangedUserId, int64(0))
	})
}
