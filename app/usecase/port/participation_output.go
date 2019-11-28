package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type ParticipationOutput interface {
	CreateUserLectureOutput
	DeleteUserLectureOutput
}

type CreateUserLectureOutput interface {
	HandleCreateStudentLecture(studentLecture *model.StudentLecture, err error) (*dto.CreateStudentLectureOutputForm, error)
	HandleCreateAssistantLecture(assistantLecture *model.AssistantLecture, err error) (*dto.CreateAssistantLectureOutputForm, error)
	HandleCreateTeacherLecture(teacherLecture *model.TeacherLecture, err error) (*dto.CreateTeacherLectureOutputForm, error)
}

type DeleteUserLectureOutput interface {
	HandleDeleteStudentLecture(id int64, err error) (*dto.DeleteStudentLectureOutputForm, error)
	HandleDeleteAssistantLecture(id int64, err error) (*dto.DeleteAssistantLectureOutputForm, error)
	HandleDeleteTeacherLecture(id int64, err error) (*dto.DeleteTeacherLectureOutputForm, error)
}
