package interactor

import (
	"fmt"
	"testing"

	"github.com/arabian9ts/sweeTest/app/interface/presenter"
	"github.com/arabian9ts/sweeTest/fixture"
	"github.com/arabian9ts/sweeTest/mock"
	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	t.Run("user=student, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		student := fixture.NewValidCreateStudentInputForm()
		outputForm, err := userInteractor.CreateStudent(student)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm)
		assert.NotEqual(t, outputForm.ID, int64(0))
	})

	t.Run("user=student, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		student := fixture.NewValidCreateStudentInputForm()
		outputForm, err := userInteractor.CreateStudent(student)

		assert.NotNil(t, err)
		assert.Equal(t, outputForm.ID, int64(0))
	})
}

func TestCreateAssistant(t *testing.T) {
	t.Run("user=assistant, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		assistant := fixture.NewValidCreateAssistantInputForm()
		outputForm, err := userInteractor.CreateAssistant(assistant)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm)
		assert.NotEqual(t, outputForm.ID, int64(0))
	})

	t.Run("user=assistant, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			AssistantFixture: fixture.NewValidAssistant(),
			Error:            fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		assistant := fixture.NewValidCreateAssistantInputForm()
		outputForm, err := userInteractor.CreateAssistant(assistant)

		assert.NotNil(t, err)
		assert.Equal(t, outputForm.ID, int64(0))
	})
}

func TestCreateTeacher(t *testing.T) {
	t.Run("user=teacher, success", func(t *testing.T) {
		userRepository := mock.NewUserRepositoryMock()
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(userRepository, userOutput)
		teacher := fixture.NewValidCreateTeacherInputForm()
		outputForm, err := userInteractor.CreateTeacher(teacher)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm)
		assert.NotEqual(t, outputForm.ID, int64(0))
	})

	t.Run("user=teacher, failed", func(t *testing.T) {
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          fmt.Errorf("mock error"),
		}
		userOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, userOutput)
		teacher := fixture.NewValidCreateTeacherInputForm()
		outputForm, err := userInteractor.CreateTeacher(teacher)

		assert.NotNil(t, err)
		assert.Equal(t, outputForm.ID, int64(0))
	})
}

func TestUpdateStudent(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		studentFixture := fixture.NewValidUpdateStudentInputForm()
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          nil,
		}
		studentOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, studentOutput)
		outputForm, err := userInteractor.UpdateStudent(studentFixture)

		assert.Nil(t, err)
		assert.Equal(t, studentFixture.ID, outputForm.ID)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		studentFixture := fixture.NewInValidUpdateStudentInputForm()
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          nil,
		}
		studentOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, studentOutput)
		outputForm, err := userInteractor.UpdateStudent(studentFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.ID, int64(0))
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		studentFixture := fixture.NewInValidUpdateStudentInputForm()
		mockRepository := &mock.UserRepositoryMock{
			StudentFixture: fixture.NewValidStudent(),
			Error:          fmt.Errorf("repository error"),
		}
		studentOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, studentOutput)
		outputForm, err := userInteractor.UpdateStudent(studentFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.ID, int64(0))
	})
}

func TestUpdateAssistant(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		assistantFixture := fixture.NewValidUpdateAssistantInputForm()
		mockRepository := &mock.UserRepositoryMock{
			AssistantFixture: fixture.NewValidAssistant(),
			Error:            nil,
		}
		assistantOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, assistantOutput)
		outputForm, err := userInteractor.UpdateAssistant(assistantFixture)

		assert.Nil(t, err)
		assert.Equal(t, assistantFixture.ID, outputForm.ID)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		assistantFixture := fixture.NewInValidUpdateAssistantInputForm()
		mockRepository := &mock.UserRepositoryMock{
			AssistantFixture: fixture.NewValidAssistant(),
			Error:            nil,
		}
		assistantOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, assistantOutput)
		assistantFixture.ID = 0
		outputForm, err := userInteractor.UpdateAssistant(assistantFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, int64(0), outputForm.ID)
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		assistantFixture := fixture.NewInValidUpdateAssistantInputForm()
		mockRepository := &mock.UserRepositoryMock{
			AssistantFixture: fixture.NewValidAssistant(),
			Error:            fmt.Errorf("repository error"),
		}
		assistantOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, assistantOutput)
		outputForm, err := userInteractor.UpdateAssistant(assistantFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, int64(0), outputForm.ID)
	})
}

func TestUpdateTeacher(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		teacherFixture := fixture.NewValidUpdateTeacherInputForm()
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          nil,
		}
		teacherOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, teacherOutput)
		outputForm, err := userInteractor.UpdateTeacher(teacherFixture)

		assert.Nil(t, err)
		assert.Equal(t, teacherFixture.ID, outputForm.ID)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		teacherFixture := fixture.NewInValidUpdateTeacherInputForm()
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          nil,
		}
		teacherOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, teacherOutput)
		teacherFixture.ID = 0
		outputForm, err := userInteractor.UpdateTeacher(teacherFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, int64(0), outputForm.ID)
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		teacherFixture := fixture.NewInValidUpdateTeacherInputForm()
		mockRepository := &mock.UserRepositoryMock{
			TeacherFixture: fixture.NewValidTeacher(),
			Error:          fmt.Errorf("repository error"),
		}
		teacherOutput := &presenter.UserPresenter{}
		userInteractor, _ := NewUserInteractor(mockRepository, teacherOutput)
		outputForm, err := userInteractor.UpdateTeacher(teacherFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, int64(0), outputForm.ID)
	})
}
