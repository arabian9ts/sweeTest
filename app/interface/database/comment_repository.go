package database

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type CommentRepository struct {
	SqlHandler
}

func NewCommentRepository(sqlHandler SqlHandler) repository.CommentRepository {
	return &CommentRepository{SqlHandler: sqlHandler}
}

func (repo *CommentRepository) GetCommentsByHelpID(helpID int64, limit int, offset int) (model.Comments, error) {
	comments := model.Comments{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `comments` WHERE `help_id` = ? LIMIT ? OFFSET ?", helpID, limit, offset)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := &model.Comment{}
		if err = rows.Scan(&comment.ID, &comment.HelpID, &comment.UserType, &comment.UserID, &comment.Contents, &comment.CreatedAt, &comment.UpdatedAt); err != nil {
			continue
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repo *CommentRepository) GetCommentsByUserTypeAndUserID(userType model.UserType, userID int64, limit int, offset int) (model.Comments, error) {
	comments := model.Comments{}
	rows, err := repo.SqlHandler.Query("SELECT * FROM `comments` WHERE `user_id` = ? AND `user_type` = ? LIMIT ? OFFSET ?", userID, userType, limit, offset)
	if err != nil {
		return comments, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := &model.Comment{}
		if err = rows.Scan(comment.ID, comment.HelpID, comment.UserType, comment.UserID, comment.Contents, comment.CreatedAt, comment.UpdatedAt); err != nil {
			continue
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repo *CommentRepository) CreateComment(comment *model.Comment) (*model.Comment, error) {
	row, err := repo.SqlHandler.Execute("INSERT INTO `comments` (`help_id`, `user_type`, `user_id`, `contents`) VALUES (?,?,?,?)",
		comment.HelpID, comment.UserType, comment.UserID, comment.Contents,
	)
	if err != nil {
		return nil, err
	}

	id, err := row.LastInsertId()
	if err != nil {
		return nil, err
	}

	comment.ID = id
	return comment, nil
}

func (repo *CommentRepository) UpdateComment(comment *model.Comment) (*model.Comment, error) {
	result, err := repo.SqlHandler.Execute("UPDATE `comments` SET `contents` = ? WHERE `id` = ? AND `user_type` = ? AND `user_id` = ?",
		comment.ID, comment.Contents, comment.UserType, comment.UserID,
	)
	if err != nil {
		return nil, err
	}

	_, err = result.RowAffected()
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (repo *CommentRepository) DeleteComment(id int64, userType model.UserType, userID int64) (int64, error) {
	result, err := repo.SqlHandler.Execute("DELETE FROM `comments` WHERE `id` = ? AND `user_type` = ? AND `user_id` = ?",
		id, userType, userID,
	)
	if err != nil {
		return 0, err
	}

	count, err := result.RowAffected()
	if err != nil {
		return 0, err
	}

	return count, nil
}
