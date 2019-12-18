package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type LectureRepository struct {
	SqlHandler
}

func NewLectureRepository(sqlHandler SqlHandler) (repository.LectureRepository, error) {
	return &LectureRepository{SqlHandler: sqlHandler}, nil
}

func (repo *LectureRepository) GetLectures(limit int, offset int) (total int64, lectures model.Lectures, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM lectures LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		lecture := model.Lecture{}
		if err = rows.Scan(&lecture.ID, &lecture.Name, &lecture.CreatedAt, &lecture.UpdatedAt); err != nil {
			continue
		}
		lectures = append(lectures, &lecture)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`id`) FROM `lectures`")
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, lectures, nil
}

func (repo *LectureRepository) GetLectureById(id int64) (lecture *model.Lecture, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM lectures WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	row.Next()

	lecture = &model.Lecture{}
	err = row.Scan(&lecture.ID, &lecture.Name, &lecture.CreatedAt, &lecture.UpdatedAt)
	return
}

func (repo *LectureRepository) GetParticipatedLecturesOfStudent(studentID int64, limit int, offset int) (total int64, lectures model.Lectures, err error) {
	rows, err := repo.SqlHandler.Query("select lectures.* from lectures inner join students_lectures on students_lectures.lecture_id = lectures.id where student_id = ? limit ? offset ?",
		studentID,
		limit,
		offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	lecture := model.Lecture{}
	for rows.Next() {
		if err = rows.Scan(&lecture.ID, &lecture.Name, &lecture.CreatedAt, &lecture.UpdatedAt); err != nil {
			continue
		}
		lectures = append(lectures, &lecture)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`student_id`) FROM `lectures` inner join `students_lectures` on students_lectures.lecture_id = lectures.id where `student_id` = ?", studentID)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, lectures, nil
}

func (repo *LectureRepository) GetParticipatedLecturesOfAssistant(studentID int64, limit int, offset int) (total int64, lectures model.Lectures, err error) {
	rows, err := repo.SqlHandler.Query("select lectures.* from lectures inner join assistants_lectures on assistants_lectures.lecture_id = lectures.id where student_id = ? limit ? offset ?",
		studentID,
		limit,
		offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		lecture := model.Lecture{}
		if err = rows.Scan(&lecture.ID, &lecture.Name, &lecture.CreatedAt, &lecture.UpdatedAt); err != nil {
			continue
		}
		lectures = append(lectures, &lecture)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`student_id`) FROM `lectures` inner join `assistants_lectures` on assistants_lectures.lecture_id = lectures.id where `student_id` = ? limit ? offset ?", studentID)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, lectures, nil
}

func (repo *LectureRepository) GetParticipatedLecturesOfTeacher(teacherID int64, limit int, offset int) (total int64, lectures model.Lectures, err error) {
	rows, err := repo.SqlHandler.Query("select lectures.* from lectures inner join teachers_lectures on teachers_lectures.lecture_id = lectures.id where teacher_id = ? limit ? offset ?",
		teacherID,
		limit,
		offset,
	)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		lecture := model.Lecture{}
		if err = rows.Scan(&lecture.ID, &lecture.Name, &lecture.CreatedAt, &lecture.UpdatedAt); err != nil {
			continue
		}
		lectures = append(lectures, &lecture)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`teacher_id`) FROM `lectures` inner join `teachers_lectures` on teachers_lectures.lecture_id = lectures.id where `teacher_id` = ?", teacherID)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, lectures, nil
}

func (repo *LectureRepository) CreateLecture(lecture *model.Lecture) (*model.Lecture, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO lectures (name) VALUES (?)",
		lecture.Name,
	)
	if err != nil {
		return nil, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	lecture.ID = id64
	return lecture, nil
}

func (repo *LectureRepository) UpdateLecture(lecture *model.Lecture) (*model.Lecture, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE lectures SET name = ? WHERE id = ?",
		lecture.Name,
		lecture.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return lecture, nil
}

func (repo *LectureRepository) DeleteLecture(id int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM lectures WHERE id = ?", id)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
