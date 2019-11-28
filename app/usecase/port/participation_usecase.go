package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type ParticipationUseCase interface {
	CreateUserLectureUseCase
	DeleteUserLectureUseCase
}

type CreateUserLectureUseCase interface {
	CreateStudentLecture(form *dto.CreateStudentLectureInputForm) (*dto.CreateStudentLectureOutputForm, error)
	CreateAssistantLecture(form *dto.CreateAssistantLectureInputForm) (*dto.CreateAssistantLectureOutputForm, error)
	CreateTeacherLecture(form *dto.CreateTeacherLectureInputForm) (*dto.CreateTeacherLectureOutputForm, error)
}

type DeleteUserLectureUseCase interface {
	DeleteStudentLecture(studentID, lectureID int64) (*dto.DeleteStudentLectureOutputForm, error)
	DeleteAssistantLecture(studentID, lectureID int64) (*dto.DeleteAssistantLectureOutputForm, error)
	DeleteTeacherLecture(studentID, lectureID int64) (*dto.DeleteTeacherLectureOutputForm, error)
}