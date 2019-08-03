package adapter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

func ConvertCreateCommentInputFormToComment(form *dto.CreateCommentInputForm) *model.Comment {
	return &model.Comment{
		HelpID:   form.HelpID,
		UserType: form.UserType,
		UserID:   form.UserID,
		Contents: form.Contents,
	}
}

func ConvertUpdateCommentInputFormToComment(form *dto.UpdateCommentInputForm) *model.Comment {
	return &model.Comment{
		ID:       form.ID,
		HelpID:   form.HelpID,
		UserType: form.UserType,
		UserID:   form.UserID,
		Contents: form.Contents,
	}
}
