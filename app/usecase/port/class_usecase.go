package port

import (
	"github.com/arabian9ts/sweeTest/app/dto"
)

type ClassUseCase interface {
	GetClassUseCase
	CreateClassUseCase
	UpdateClassUseCase
	DeleteClassUseCase
}

type GetClassUseCase interface {
	GetClassById(id int64) (*dto.GetClassByIdOutputForm, error)
	GetClassesByLectureId(lectureID int64, limit int, offset int) (dto.GetClassesOutputForm, error)
}

type CreateClassUseCase interface {
	CreateClass(form *dto.CreateClassInputForm) (*dto.CreateClassOutputForm, error)
}

type UpdateClassUseCase interface {
	UpdateClass(form *dto.UpdateClassInputForm) (*dto.UpdateClassOutputForm, error)
}

type DeleteClassUseCase interface {
	DeleteClass(id int64, lectureID int64) (*dto.DeleteClassOutputForm, error)
}
