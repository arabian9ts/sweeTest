package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type HelpOutput interface {
	GetHelpOutput
	CreateHelpOutput
	UpdateHelpOutput
	DeleteHelpOutput
}

type GetHelpOutput interface {
	HandleGetHelpsByLectureID(total int64, helps model.Helps, err error) (dto.GetTotalHelpsOutputForm, error)
}

type CreateHelpOutput interface {
	HandleCreateHelp(help *model.Help, err error) (*dto.CreateHelpOutputForm, error)
}

type UpdateHelpOutput interface {
	HandleUpdateHelp(help *model.Help, err error) (*dto.UpdateHelpOutputForm, error)
}

type DeleteHelpOutput interface {
	HandleDeleteHelp(id int64, err error) (*dto.DeleteHelpOutputForm, error)
}
