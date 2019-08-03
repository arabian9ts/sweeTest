package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type CommentInteractor struct {
	CommentRepository repository.CommentRepository
	Output            port.CommentOutput
}

func NewCommentInteractor(repository repository.CommentRepository, output port.CommentOutput) (*CommentInteractor, error) {
	return &CommentInteractor{
		CommentRepository: repository,
		Output:            output,
	}, nil
}

func (interactor *CommentInteractor) GetCommentsByHelpID(helpID int64, limit int, offset int) (dto.GetCommentsOutputForm, error) {
	return interactor.Output.HandleGetComments(
		interactor.CommentRepository.GetCommentsByHelpID(helpID, limit, offset),
	)
}

func (interactor *CommentInteractor) CreateComment(form *dto.CreateCommentInputForm) (*dto.CreateCommentOutputForm, error) {
	comment := adapter.ConvertCreateCommentInputFormToComment(form)
	return interactor.Output.HandleCreateComment(
		interactor.CommentRepository.CreateComment(comment),
	)
}

func (interactor *CommentInteractor) UpdateComment(form *dto.UpdateCommentInputForm) (*dto.UpdateCommentOutputForm, error) {
	comment := adapter.ConvertUpdateCommentInputFormToComment(form)
	return interactor.Output.HandleUpdateComment(
		interactor.CommentRepository.CreateComment(comment),
	)
}

func (interactor *CommentInteractor) DeleteComment(id int64, userType model.UserType, userID int64) (*dto.DeleteCommentOutputForm, error) {
	return interactor.Output.HandleDeleteComment(
		interactor.CommentRepository.DeleteComment(id, userType, userID),
	)
}
