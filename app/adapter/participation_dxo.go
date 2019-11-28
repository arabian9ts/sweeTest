package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateStudentLectureInputFormToStudentLecture(form *dto.CreateStudentLectureInputForm) (studentLecture *model.StudentLecture) {
	studentLecture = &model.StudentLecture{
		StudentID: form.StudentID,
		LectureID: form.LectureID,
	}
	return
}

func ConvertCreateAssistantLectureInputFormToAssistantLecture(form *dto.CreateAssistantLectureInputForm) (assistantLecture *model.AssistantLecture) {
	assistantLecture = &model.AssistantLecture{
		StudentID: form.StudentID,
		LectureID: form.LectureID,
	}
	return
}

func ConvertCreateTeacherLectureInputFormToTeacherLecture(form *dto.CreateTeacherLectureInputForm) (teacherLecture *model.TeacherLecture) {
	teacherLecture = &model.TeacherLecture{
		TeacherID: form.TeacherID,
		LectureID: form.LectureID,
	}
	return
}
