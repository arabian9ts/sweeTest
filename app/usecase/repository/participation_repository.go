package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type ParticipationRepository interface {
	CreateStudentLecture(studentLecture *model.StudentLecture) (*model.StudentLecture, error)
	CreateAssistantLecture(assistantLecture *model.AssistantLecture) (*model.AssistantLecture, error)
	CreateTeacherLecture(teacherLecture *model.TeacherLecture) (*model.TeacherLecture, error)

	DeleteStudentLecture(studentID int64, lectureID int64) (int64, error)
	DeleteAssistantLecture(studentID int64, lectureID int64) (int64, error)
	DeleteTeacherLecture(teacherID int64, lectureID int64) (int64, error)
}
