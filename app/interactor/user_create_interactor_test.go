package interactor

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/fixture"
	"github.com/arabian9ts/sweeTest/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateStudent(t *testing.T) {
	t.Run("user=student, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userInteractor, _ := NewCreateUserInteractor(userRepository)
		inputForm := fixture.NewCreateStudentInputFormMock()
		outputForm := userInteractor.CreateStudent(inputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
		assert.NotEqual(t, outputForm.JwtToken, "")
	})

	t.Run("user=student, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          fmt.Errorf("mock error"),
		}
		userInteractor, _ := NewCreateUserInteractor(mockRepository)
		inputForm := fixture.NewCreateStudentInputFormMock()
		outputForm := userInteractor.CreateStudent(inputForm)
		fmt.Println(outputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
		assert.Equal(t, outputForm.JwtToken, "")
	})
}

func TestCreateTa(t *testing.T) {
	t.Run("user=ta, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userInteractor, _ := NewCreateUserInteractor(userRepository)
		inputForm := fixture.NewCreateTaInputFormMock()
		outputForm := userInteractor.CreateTa(inputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
		assert.NotEqual(t, outputForm.JwtToken, "")
	})

	t.Run("user=ta, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TaFixture: fixture.NewValidTa(),
			Error:     fmt.Errorf("mock error"),
		}
		userInteractor, _ := NewCreateUserInteractor(mockRepository)
		inputForm := fixture.NewCreateTaInputFormMock()
		outputForm := userInteractor.CreateTa(inputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
		assert.Equal(t, outputForm.JwtToken, "")
	})
}

func TestCreateTeacher(t *testing.T) {
	t.Run("user=teacher, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userInteractor, _ := NewCreateUserInteractor(userRepository)
		inputForm := fixture.NewCreateTeacherInputFormMock()
		outputForm := userInteractor.CreateTeacher(inputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.NotEqual(t, outputForm.LastCreatedUserId, 0)
		assert.NotEqual(t, outputForm.JwtToken, "")
	})

	t.Run("user=teacher, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          fmt.Errorf("mock error"),
		}
		userInteractor, _ := NewCreateUserInteractor(mockRepository)
		inputForm := fixture.NewCreateTeacherInputFormMock()
		outputForm := userInteractor.CreateTeacher(inputForm)
		fmt.Println(outputForm)

		assert.NotNil(t, outputForm.LastCreatedUserId)
		assert.NotNil(t, outputForm.JwtToken)
		assert.Equal(t, outputForm.LastCreatedUserId, int64(0))
		assert.Equal(t, outputForm.JwtToken, "")
	})
}
