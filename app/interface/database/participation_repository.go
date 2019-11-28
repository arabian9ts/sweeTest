package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type ParticipationRepository struct {
	SqlHandler
}

func NewParticipationRepository(sqlHandler SqlHandler) (repository.ParticipationRepository, error) {
	return &ParticipationRepository{SqlHandler: sqlHandler}, nil
}

func (repo *ParticipationRepository) CreateStudentLecture(studentLecture *model.StudentLecture) (*model.StudentLecture, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `students_lectures` (`student_id`, `lecture_id`) VALUES (?,?)",
		studentLecture.StudentID,
		studentLecture.LectureID,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	studentLecture.ID = id64
	return studentLecture, nil
}

func (repo *ParticipationRepository) DeleteStudentLecture(studentID int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"DELETE FROM students_lectures WHERE student_id = ? AND lecture_id = ?",
		studentID,
		lectureID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}

func (repo *ParticipationRepository) CreateAssistantLecture(assistantLecture *model.AssistantLecture) (*model.AssistantLecture, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `assistants_lectures` (`student_id`, `lecture_id`) VALUES (?,?)",
		assistantLecture.StudentID,
		assistantLecture.LectureID,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	assistantLecture.ID = id64
	return assistantLecture, nil
}

func (repo *ParticipationRepository) DeleteAssistantLecture(studentID int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"DELETE FROM assistants_lectures WHERE student_id = ? AND lecture_id = ?",
		studentID,
		lectureID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}

func (repo *ParticipationRepository) CreateTeacherLecture(teacherLecture *model.TeacherLecture) (*model.TeacherLecture, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `teachers_lectures` (`teacher_id`, `lecture_id`) VALUES (?,?)",
		teacherLecture.TeacherID,
		teacherLecture.LectureID,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	teacherLecture.ID = id64
	return teacherLecture, nil
}

func (repo *ParticipationRepository) DeleteTeacherLecture(teacherID int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"DELETE FROM teachers_lectures WHERE teacher_id = ? AND lecture_id = ?",
		teacherID,
		lectureID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
