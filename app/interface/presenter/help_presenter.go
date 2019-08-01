package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type HelpPresenter struct{}

func NewHelpPresenter() port.HelpOutput {
	return &HelpPresenter{}
}

func (*HelpPresenter) HandleGetHelpsByLectureID(helps model.Helps, err error) (dto.GetHelpsOutputForm, error) {
	output := dto.GetHelpsOutputForm{}
	for _, help := range helps {
		form := &dto.GetHelpOutputForm{
			ID:        help.ID,
			LectureID: help.LectureID,
			StudentID: help.StudentID,
			Contents:  help.Contents,
			CreatedAt: help.CreatedAt,
			UpdatedAt: help.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (*HelpPresenter) HandleCreateHelp(id int64, err error) (*dto.CreateHelpOutputForm, error) {
	return &dto.CreateHelpOutputForm{
		LastChangedHelpID: id,
	}, err
}

func (*HelpPresenter) HandleUpdateHelp(count int64, err error) (*dto.UpdateHelpOutputForm, error) {
	updated := count != 0
	return &dto.UpdateHelpOutputForm{
		Updated: updated,
	}, err
}

func (*HelpPresenter) HandleDeleteHelp(count int64, err error) (*dto.DeleteHelpOutputForm, error) {
	deleted := count != 0
	return &dto.DeleteHelpOutputForm{
		Deleted: deleted,
	}, err
}
