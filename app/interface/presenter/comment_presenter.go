package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type CommentPresenter struct{}

func NewCommentPresenter() port.CommentOutput {
	return &CommentPresenter{}
}

func (*CommentPresenter) HandleGetComments(comments model.Comments, err error) (dto.GetCommentsOutputForm, error) {
	output := dto.GetCommentsOutputForm{}
	for _, comment := range comments {
		form := &dto.GetCommentOutputForm{
			ID:        comment.ID,
			HelpID:    comment.HelpID,
			UserType:  model.UserType.String(comment.UserType),
			UserID:    comment.UserID,
			Contents:  comment.Contents,
			CreatedAt: comment.CreatedAt,
			UpdatedAt: comment.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*CommentPresenter) HandleCreateComment(id int64, err error) (*dto.CreateCommentOutputForm, error) {
	return &dto.CreateCommentOutputForm{
		ID: id,
	}, err
}

func (*CommentPresenter) HandleUpdateComment(count int64, err error) (*dto.UpdateCommentOutputForm, error) {
	updated := count != 0
	return &dto.UpdateCommentOutputForm{
		Updated: updated,
	}, err
}

func (*CommentPresenter) HandleDeleteComment(count int64, err error) (*dto.DeleteCommentOutputForm, error) {
	deleted := count != 0
	return &dto.DeleteCommentOutputForm{
		Deleted: deleted,
	}, err
}
