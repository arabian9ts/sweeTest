package database

import (
	"fmt"
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type HelpRepository struct {
	SqlHandler
}

func NewHelpRepository(sqlHandler SqlHandler) (repository.HelpRepository, error) {
	return &HelpRepository{SqlHandler: sqlHandler}, nil
}

func (repo *HelpRepository) GetHelpsByStudentID(studentID int64, limit int, offset int) (model.Helps, error) {
	helps := model.Helps{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `helps` WHERE `student_id` = ? LIMIT ? OFFSET ?", studentID, limit, offset)
	if err != nil {
		return helps, err
	}
	defer rows.Close()

	for rows.Next() {
		help := &model.Help{}
		if err = rows.Scan(help.ID, help.LectureID, help.StudentID, help.Contents, help.CreatedAt, help.UpdatedAt); err != nil {
			continue
		}
		helps = append(helps, help)
	}
	return helps, nil
}

func (repo *HelpRepository) GetHelpsByLectureID(lectureID int64, limit int, offset int) (model.Helps, error) {
	helps := model.Helps{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `helps` WHERE `lecture_id` = ? LIMIT ? OFFSET ?", lectureID, limit, offset)
	if err != nil {
		return helps, err
	}
	defer rows.Close()

	for rows.Next() {
		help := &model.Help{}
		if err = rows.Scan(&help.ID, &help.LectureID, &help.StudentID, &help.Contents, &help.CreatedAt, &help.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		helps = append(helps, help)
	}
	return helps, nil
}

func (repo *HelpRepository) CreateHelp(help *model.Help) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `helps` (`lecture_id`, `student_id`, `contents`) VALUES (?,?,?)",
		help.LectureID, help.StudentID, help.Contents,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int64(count), err
}

func (repo *HelpRepository) UpdateHelp(help *model.Help) (int64, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE `helps` SET `contents` = ? WHERE `lecture_id` = ? AND `student_id` = ?",
		help.Contents, help.LectureID, help.StudentID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), err
}

func (repo *HelpRepository) DeleteHelp(id int64, lectureID int64, studentID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `helps` WHERE `id` = ? AND `lecture_id` = ? AND `student_id` = ?", id, lectureID, studentID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return int64(count), nil
}
