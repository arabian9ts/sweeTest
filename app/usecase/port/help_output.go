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
	HandleGetHelpsByLectureID(helps model.Helps, err error) (dto.GetHelpsOutputForm, error)
}

type CreateHelpOutput interface {
	HandleCreateHelp(id int64, err error) (*dto.CreateHelpOutputForm, error)
}

type UpdateHelpOutput interface {
	HandleUpdateHelp(id int64, err error) (*dto.UpdateHelpOutputForm, error)
}

type DeleteHelpOutput interface {
	HandleDeleteHelp(id int64, err error) (*dto.DeleteHelpOutputForm, error)
}
