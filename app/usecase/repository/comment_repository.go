package repository

import "github.com/arabian9ts/sweeTest/app/domain/model"

type CommentRepository interface {
	GetCommentsByHelpID(helpID int64, limit int, offset int) (model.Comments, error)
	GetCommentsByUserTypeAndUserID(userType model.UserType, userID int64, limit int, offset int) (model.Comments, error)
	CreateComment(comment *model.Comment) (int64, error)
	UpdateComment(comment *model.Comment) (int64, error)
	DeleteComment(id int64, userType model.UserType, userID int64) (int64, error)
}
