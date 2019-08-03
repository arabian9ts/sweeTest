package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type CommentUseCase interface {
	GetCommentUseCase
	CreateCommentUseCase
	UpdateCommentUseCase
	DeleteCommentUseCase
}

type GetCommentUseCase interface {
	GetCommentsByHelpID(helpID int64, limit int, offset int) (dto.GetCommentsOutputForm, error)
}

type CreateCommentUseCase interface {
	CreateComment(form *dto.CreateCommentInputForm) (*dto.CreateCommentOutputForm, error)
}

type UpdateCommentUseCase interface {
	UpdateComment(form *dto.UpdateCommentInputForm) (*dto.UpdateCommentOutputForm, error)
}

type DeleteCommentUseCase interface {
	DeleteComment(id int64, userType model.UserType, userID int64) (*dto.DeleteCommentOutputForm, error)
}
