package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type CommentOutput interface {
	GetCommentOutput
	CreateCommentOutput
	UpdateCommentOutput
	DeleteCommentOutput
}

type GetCommentOutput interface {
	HandleGetComments(comments model.Comments, err error) (dto.GetCommentsOutputForm, error)
}

type CreateCommentOutput interface {
	HandleCreateComment(comment *model.Comment, err error) (*dto.CreateCommentOutputForm, error)
}

type UpdateCommentOutput interface {
	HandleUpdateComment(comment *model.Comment, err error) (*dto.UpdateCommentOutputForm, error)
}

type DeleteCommentOutput interface {
	HandleDeleteComment(count int64, err error) (*dto.DeleteCommentOutputForm, error)
}
