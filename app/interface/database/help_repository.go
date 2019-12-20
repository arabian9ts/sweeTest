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

func (repo *HelpRepository) GetHelpsByStudentID(studentID int64, limit int, offset int) (total int64, helps model.Helps, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `helps` WHERE `student_id` = ? LIMIT ? OFFSET ?", studentID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		help := &model.Help{}
		if err = rows.Scan(help.ID, help.LectureID, help.StudentID, help.Title, help.Contents, help.CreatedAt, help.UpdatedAt); err != nil {
			continue
		}
		helps = append(helps, help)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`id`) FROM `helps` WHERE `student_id` = ?", studentID)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, helps, nil
}

func (repo *HelpRepository) GetHelpsByLectureID(lectureID int64, limit int, offset int) (total int64, helps model.Helps, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `helps` WHERE `lecture_id` = ? LIMIT ? OFFSET ?", lectureID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		help := &model.Help{}
		if err = rows.Scan(&help.ID, &help.LectureID, &help.StudentID, &help.Title, &help.Contents, &help.CreatedAt, &help.UpdatedAt); err != nil {
			fmt.Println(err)
			continue
		}
		helps = append(helps, help)
	}

	row, err := repo.SqlHandler.Query("SELECT COUNT(`id`) FROM `helps` WHERE `lecture_id` = ?", lectureID)
	if err != nil {
		return
	}
	defer row.Close()
	row.Next()

	err = row.Scan(&total)
	return total, helps, nil
}

func (repo *HelpRepository) CreateHelp(help *model.Help) (*model.Help, error) {
	result, err := repo.SqlHandler.Execute(
		"INSERT INTO `helps` (`lecture_id`, `student_id`, `title`, `contents`) VALUES (?,?,?,?)",
		help.LectureID, help.StudentID, help.Title, help.Contents,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	help.ID = id
	return help, err
}

func (repo *HelpRepository) UpdateHelp(help *model.Help) (*model.Help, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE `helps` SET `title` = ?, `contents` = ? WHERE `lecture_id` = ? AND `student_id` = ?",
		help.Title, help.Contents, help.LectureID, help.StudentID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return help, err
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

	return count, nil
}

func (repo *HelpRepository) DeleteHelpWithoutStudentId(id int64, lectureID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `helps` WHERE `id` = ? AND `lecture_id` = ?", id, lectureID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
