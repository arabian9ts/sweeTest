package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type SubmissionRepository struct {
	SqlHandler
}

func NewSubmissionRepository(sqlHandler SqlHandler) (repository.SubmissionRepository, error) {
	return &SubmissionRepository{SqlHandler: sqlHandler}, nil
}

func (repo *SubmissionRepository) GetSubmissionTextByID(submissionTextID int64) (submissionText *model.SubmissionText, err error) {
	row, err := repo.SqlHandler.Query("SELECT * FROM `submission_texts` WHERE `id` = ?", submissionTextID)
	if err != nil {
		return &model.SubmissionText{}, err
	}
	defer row.Close()
	row.Next()

	submissionText = &model.SubmissionText{}
	err = row.Scan(&submissionText.ID, &submissionText.FileName, &submissionText.Contents, &submissionText.CreatedAt, &submissionText.UpdatedAt)
	return
}

func (repo *SubmissionRepository) GetSubmissionTextsBySubmissionID(submissionID int64, limit int, offset int) (submissionTexts model.SubmissionTexts, err error) {
	rows, err := repo.SqlHandler.Query("SELECT * FROM `submission_texts` WHERE `submission_id` = ? LIMIT ? OFFSET ?", submissionID, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		submissionText := &model.SubmissionText{}
		if err = rows.Scan(submissionText.ID, submissionText.SubmissionID, submissionText.FileName, submissionText.Contents, submissionText.CreatedAt, submissionText.UpdatedAt); err != nil {
			continue
		}
		submissionTexts = append(submissionTexts, submissionText)
	}
	return submissionTexts, nil
}

func (repo *SubmissionRepository) CreateSubmission(submission *model.Submission) (*model.Submission, error) {
	submissionResult, err := repo.SqlHandler.Execute(
		"INSERT INTO `submissions` (`task_id`, `student_id`) VALUES (?,?)",
		submission.TaskID, submission.StudentID,
	)
	if err != nil {
		return nil, err
	}

	submissionID, err := submissionResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	submission.ID = submissionID
	return submission, err
}

func (repo *SubmissionRepository) CreateSubmissionText(submissionText *model.SubmissionText, submissionID int64) (*model.SubmissionText, error) {
	submissionTextResult, err := repo.SqlHandler.Execute(
		"INSERT INTO `submission_texts` (`submission_id`, `file_name`, `contents`) VALUES (?,?,?)",
		submissionID, submissionText.FileName, submissionText.Contents,
	)
	if err != nil {
		return nil, err
	}

	submissionTextID, err := submissionTextResult.LastInsertId()
	if err != nil {
		return nil, err
	}

	submissionText.ID = submissionTextID
	return submissionText, err
}

func (repo *SubmissionRepository) UpdateSubmissionText(submissionText *model.SubmissionText) (*model.SubmissionText, error) {
	result, err := repo.SqlHandler.Execute(
		"UPDATE `submission_texts` SET `file_name` = ?, `contents` = ? WHERE `id` = ?",
		submissionText.FileName, submissionText.Contents, submissionText.ID,
	)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	submissionText.ID = id
	return submissionText, err
}

func (repo *SubmissionRepository) DeleteSubmissionText(id int64, submissionID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `submission_texts` WHERE `id` = ? AND `submission_id` = ?", id, submissionID)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
