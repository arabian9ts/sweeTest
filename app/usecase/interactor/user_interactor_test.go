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
		inputForm := fixture.NewCreateStudentInputFormMock()
		outputForm, err := userInteractor.CreateStudent(inputForm)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
	})

	t.Run("user=student, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		inputForm := fixture.NewCreateStudentInputFormMock()
		outputForm, err := userInteractor.CreateStudent(inputForm)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
	})
}

func TestCreateTa(t *testing.T) {
	t.Run("user=ta, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		inputForm := fixture.NewCreateTaInputFormMock()
		outputForm, err := userInteractor.CreateTa(inputForm)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
	})

	t.Run("user=ta, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TaFixture: fixture.NewValidTa(),
			Error:     fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		inputForm := fixture.NewCreateTaInputFormMock()
		outputForm, err := userInteractor.CreateTa(inputForm)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
	})
}

func TestCreateTeacher(t *testing.T) {
	t.Run("user=teacher, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		inputForm := fixture.NewCreateTeacherInputFormMock()
		outputForm, err := userInteractor.CreateTeacher(inputForm)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
	})

	t.Run("user=teacher, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		inputForm := fixture.NewCreateTeacherInputFormMock()
		outputForm, err := userInteractor.CreateTeacher(inputForm)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
	})
}
