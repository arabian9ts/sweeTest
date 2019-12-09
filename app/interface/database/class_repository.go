package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type ClassRepository struct {
	SqlHandler
}

func NewClassRepository(sqlHandler SqlHandler) (repository.ClassRepository, error) {
	return &ClassRepository{SqlHandler: sqlHandler}, nil
}

func (repo *ClassRepository) GetClassesByLectureID(lectureID int64, limit int, offset int) (model.Classes, error) {
	classes := model.Classes{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `classes` WHERE `lecture_id` = ? LIMIT ? OFFSET ?", lectureID, limit, offset)
	if err != nil {
		return classes, err
	}
	defer rows.Close()

	for rows.Next() {
		class := &model.Class{}
		if err = rows.Scan(&class.ID, &class.LectureID, &class.Name, &class.CreatedAt, &class.UpdatedAt); err != nil {
			continue
		}
		classes = append(classes, class)
	}
	return classes, nil
}

func (repo *ClassRepository) GetClassById(id int64) (class *model.Class, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM classes WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()
	row.Next()

	class = &model.Class{}
	err = row.Scan(&class.ID, &class.LectureID, &class.Name, &class.CreatedAt, &class.UpdatedAt)
	return
}

func (repo *ClassRepository) CreateClass(class *model.Class) (*model.Class, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `classes` (`lecture_id`, `name`) VALUES (?,?)",
		class.LectureID, class.Name,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	class.ID = id
	return class, err
}

func (repo *ClassRepository) UpdateClass(class *model.Class) (*model.Class, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE `classes` SET `name` = ? WHERE `id` = ?",
		class.Name, class.ID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return class, err
}

func (repo *ClassRepository) DeleteClass(id int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `classes` WHERE `id` = ? AND `lecture_id` = ?", id, lectureID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
