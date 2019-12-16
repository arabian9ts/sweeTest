package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type HelpUseCase interface {
	GetHelpUseCase
	CreateHelpUseCase
	UpdateHelpUseCase
	DeleteHelpUseCase
}

type GetHelpUseCase interface {
	GetHelpsByLectureID(lectureID int64, limit int, offset int) (dto.GetHelpsOutputForm, error)
}

type CreateHelpUseCase interface {
	CreateHelp(form *dto.CreateHelpInputForm) (*dto.CreateHelpOutputForm, error)
}

type UpdateHelpUseCase interface {
	UpdateHelp(form *dto.UpdateHelpInputForm) (*dto.UpdateHelpOutputForm, error)
}

type DeleteHelpUseCase interface {
	DeleteHelp(id int64, lectureID int64, studentID int64) (*dto.DeleteHelpOutputForm, error)
	DeleteHelpWithoutStudentId(id int64, lectureID int64) (*dto.DeleteHelpOutputForm, error)
}
