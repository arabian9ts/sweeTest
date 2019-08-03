package dto

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"time"
)

// Input
type CreateCommentInputForm struct {
	HelpID   int64          `validate:"required,gt=0" json:"-"`
	UserType model.UserType `validate:"lte=2" json:"-"`
	UserID   int64          `validate:"required,gt=0" json:"-"`
	Contents string         `validate:"required" json:"contents"`
}

type UpdateCommentInputForm struct {
	ID       int64          `validate:"required,gt=0" json:"-"`
	HelpID   int64          `validate:"required,gt=0" json:"-"`
	UserType model.UserType `validate:"lte=2" json:"-"`
	UserID   int64          `validate:"required,gt=0" json:"-"`
	Contents string         `validate:"required" json:"contents"`
}

// Output
type GetCommentOutputForm struct {
	ID        int64     `json:"id"`
	HelpID    int64     `json:"help_id"`
	UserType  string    `json:"user_type"`
	UserID    int64     `json:"user_id"`
	Contents  string    `json:"contents"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type GetCommentsOutputForm []*GetCommentOutputForm

type CreateCommentOutputForm struct {
	ID int64 `json:"id"`
}

type UpdateCommentOutputForm struct {
	Updated bool `json:"updated"`
}

type DeleteCommentOutputForm struct {
	Deleted bool `json:"deleted"`
}
