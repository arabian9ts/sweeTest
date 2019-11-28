package mock

import (
	"fmt"

	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/fixture"
)

type LectureRepositoryMock struct {
	LectureFixture *model.Lecture
	Error          error
}

func NewLectureRepositoryMock() repository.LectureRepository {
	return &LectureRepositoryMock{
		LectureFixture: fixture.NewValidLecture(),
		Error:          nil,
	}
}

func (mock *LectureRepositoryMock) GetLectures(limit int, offset int) (model.Lectures, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Lecture{mock.LectureFixture}, nil
}

func (mock *LectureRepositoryMock) GetLectureById(id int64) (*model.Lecture, error) {
	if mock.Error != nil {
		return &model.Lecture{}, mock.Error
	}
	if id != mock.LectureFixture.ID {
		return &model.Lecture{}, fmt.Errorf("lecture not exists")
	}
	return mock.LectureFixture, nil
}

func (mock *LectureRepositoryMock) GetParticipatedLecturesOfStudent(studentID int64, limit int, offset int) (model.Lectures, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Lecture{mock.LectureFixture}, nil
}

func (mock *LectureRepositoryMock) GetParticipatedLecturesOfAssistant(studentID int64, limit int, offset int) (model.Lectures, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Lecture{mock.LectureFixture}, nil
}

func (mock *LectureRepositoryMock) GetParticipatedLecturesOfTeacher(teacherID int64, limit int, offset int) (model.Lectures, error) {
	if mock.Error != nil {
		return nil, mock.Error
	}
	return []*model.Lecture{mock.LectureFixture}, nil
}

func (mock *LectureRepositoryMock) CreateLecture(lecture *model.Lecture) (*model.Lecture, error) {
	if mock.Error != nil {
		return &model.Lecture{}, mock.Error
	}
	return mock.LectureFixture, nil
}

func (mock *LectureRepositoryMock) UpdateLecture(lecture *model.Lecture) (*model.Lecture, error) {
	if mock.Error != nil {
		return &model.Lecture{}, mock.Error
	}
	if lecture.ID != mock.LectureFixture.ID {
		return &model.Lecture{}, fmt.Errorf("lecture not exists")
	}
	return lecture, nil
}

func (mock *LectureRepositoryMock) DeleteLecture(id int64) (int64, error) {
	if mock.Error != nil {
		return 0, mock.Error
	}
	if id != mock.LectureFixture.ID {
		return int64(0), fmt.Errorf("lecture not exists")
	}
	return mock.LectureFixture.ID, nil
}
