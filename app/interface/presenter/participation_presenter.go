package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type ParticipationPresenter struct {
}

func NewParticipationPresenter() port.ParticipationOutput {
	return &ParticipationPresenter{}
}

func (p ParticipationPresenter) HandleCreateStudentLecture(studentLecture *model.StudentLecture, err error) (*dto.CreateStudentLectureOutputForm, error) {
	output := &dto.CreateStudentLectureOutputForm{
		ID:        studentLecture.ID,
		StudentID: studentLecture.StudentID,
		LectureID: studentLecture.LectureID,
		CreatedAt: studentLecture.CreatedAt,
		UpdatedAt: studentLecture.UpdatedAt,
	}
	return output, err
}

func (p ParticipationPresenter) HandleCreateAssistantLecture(assistantLecture *model.AssistantLecture, err error) (*dto.CreateAssistantLectureOutputForm, error) {
	output := &dto.CreateAssistantLectureOutputForm{
		ID:        assistantLecture.ID,
		StudentID: assistantLecture.StudentID,
		LectureID: assistantLecture.LectureID,
		CreatedAt: assistantLecture.CreatedAt,
		UpdatedAt: assistantLecture.UpdatedAt,
	}
	return output, err
}

func (p ParticipationPresenter) HandleCreateTeacherLecture(teacherLecture *model.TeacherLecture, err error) (*dto.CreateTeacherLectureOutputForm, error) {
	output := &dto.CreateTeacherLectureOutputForm{
		ID:        teacherLecture.ID,
		TeacherID: teacherLecture.TeacherID,
		LectureID: teacherLecture.LectureID,
		CreatedAt: teacherLecture.CreatedAt,
		UpdatedAt: teacherLecture.UpdatedAt,
	}
	return output, err
}

func (p ParticipationPresenter) HandleDeleteStudentLecture(count int64, err error) (*dto.DeleteStudentLectureOutputForm, error) {
	output := &dto.DeleteStudentLectureOutputForm{AffectedRowsCount: count}

	return output, err
}

func (p ParticipationPresenter) HandleDeleteAssistantLecture(count int64, err error) (*dto.DeleteAssistantLectureOutputForm, error) {
	output := &dto.DeleteAssistantLectureOutputForm{AffectedRowsCount: count}

	return output, err
}

func (p ParticipationPresenter) HandleDeleteTeacherLecture(count int64, err error) (*dto.DeleteTeacherLectureOutputForm, error) {
	output := &dto.DeleteTeacherLectureOutputForm{AffectedRowsCount: count}

	return output, err
}
