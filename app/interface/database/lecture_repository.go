package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"time"
)

type LectureRepository struct {
	SqlHandler
}

func NewLectureRepository(sqlHandler SqlHandler) (repository.LectureRepository, error) {
	return &LectureRepository{SqlHandler: sqlHandler}, nil
}

func (repo *LectureRepository) GetLectures(limit int, offset int) (model.Lectures, error) {
	lectures := model.Lectures{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM lectures LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return lectures, err
	}
	defer rows.Close()

	var id int64
	var name string
	var createdAt time.Time
	var updatedAt time.Time
	for rows.Next() {
		if err = rows.Scan(&id, &name, &createdAt, &updatedAt); err != nil {
			continue
		}
		lectures = append(lectures, &model.Lecture{
			ID:        id,
			Name:      name,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}
	return lectures, nil
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
