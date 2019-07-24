package interactor

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/interface/presenter"
	"github.com/arabian9ts/sweeTest/fixture"
	"github.com/arabian9ts/sweeTest/mock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetLectures(t *testing.T) {
	t.Run("lectures.size=1, success", func(t *testing.T) {
		lectureRepository := mock.NewLectureRepositoryMock()
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(lectureRepository, lectureOutput)
		outputForm, err := lectureInteractor.GetLectures(1, 0)

		assert.Nil(t, err)
		assert.NotEqual(t, len(outputForm), 0)
	})

	t.Run("lectures.size=0, failed", func(t *testing.T) {
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: nil,
			Error:          fmt.Errorf("mock error"),
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.GetLectures(0, 0)

		assert.NotNil(t, err)
		assert.Equal(t, len(outputForm), 0)
	})
}

func TestGetLectureById(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		fixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture,
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.GetLectureById(fixture.ID)
		expected := &dto.GetLectureByIdOutputForm{
			ID:        fixture.ID,
			Name:      fixture.Name,
			CreatedAt: fixture.CreatedAt,
			UpdatedAt: fixture.UpdatedAt,
		}

		assert.Nil(t, err)
		assert.Equal(t, outputForm, expected)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.GetLectureById(0)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.ID, int64(0))
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          fmt.Errorf("repository error"),
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.GetLectureById(0)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.ID, int64(0))
	})
}

func TestCreateLecture(t *testing.T) {
	t.Run("err=nil, success", func(t *testing.T) {
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		lecture := &model.Lecture{Name: "test lecture interactor"}
		outputForm, err := lectureInteractor.CreateLecture(lecture)

		assert.Nil(t, err)
		assert.NotNil(t, outputForm)
		assert.NotEqual(t, outputForm.LastChangedLectureId, int64(0))
	})

	t.Run("err=not nil, failed", func(t *testing.T) {
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          fmt.Errorf("repository error"),
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		lecture := &model.Lecture{Name: "test lecture interactor"}
		outputForm, err := lectureInteractor.CreateLecture(lecture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.LastChangedLectureId, int64(0))
	})
}

func TestDeleteLecture(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: lectureFixture,
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.DeleteLecture(lectureFixture.ID)
		expected := &dto.DeleteLectureOutputForm{Deleted: true}

		assert.Nil(t, err)
		assert.Equal(t, outputForm, expected)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		lectureFixture.ID = 0
		outputForm, err := lectureInteractor.DeleteLecture(lectureFixture.ID)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.Deleted, false)
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          fmt.Errorf("repository error"),
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.DeleteLecture(lectureFixture.ID)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.Deleted, false)
	})
}

func TestUpdateLecture(t *testing.T) {
	t.Run("id=1, err=nil, success", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: lectureFixture,
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.UpdateLecture(lectureFixture)
		expected := &dto.UpdateLectureOutputForm{Updated: true}

		assert.Nil(t, err)
		assert.Equal(t, outputForm, expected)
	})

	t.Run("id=0, err=nil, failed", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          nil,
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		lectureFixture.ID = 0
		outputForm, err := lectureInteractor.UpdateLecture(lectureFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.Updated, false)
	})

	t.Run("id=1, err=not nil, failed", func(t *testing.T) {
		lectureFixture := fixture.NewValidLecture()
		mockRepository := &mock.LectureRepositoryMock{
			LectureFixture: fixture.NewValidLecture(),
			Error:          fmt.Errorf("repository error"),
		}
		lectureOutput := &presenter.LecturePresenter{}
		lectureInteractor, _ := NewLectureInteractor(mockRepository, lectureOutput)
		outputForm, err := lectureInteractor.UpdateLecture(lectureFixture)

		assert.NotNil(t, err)
		assert.NotNil(t, outputForm)
		assert.Equal(t, outputForm.Updated, false)
	})
}
