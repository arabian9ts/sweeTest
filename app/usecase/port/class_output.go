package port

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
)

type ClassOutput interface {
	GetClassOutput
	CreateClassOutput
	UpdateClassOutput
	DeleteClassOutput
}

type GetClassOutput interface {
	HandleGetClasses(classes model.Classes, err error) (dto.GetClassesOutputForm, error)
	HandleGetClassById(class *model.Class, err error) (*dto.GetClassByIdOutputForm, error)
}

type CreateClassOutput interface {
	HandleCreateClass(class *model.Class, err error) (*dto.CreateClassOutputForm, error)
}

type UpdateClassOutput interface {
	HandleUpdateClass(class *model.Class, err error) (*dto.UpdateClassOutputForm, error)
}

type DeleteClassOutput interface {
	HandleDeleteClass(count int64, err error) (*dto.DeleteClassOutputForm, error)
}
