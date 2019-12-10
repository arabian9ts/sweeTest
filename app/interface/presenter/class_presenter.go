package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type ClassPresenter struct {
}

func NewClassPresenter() port.ClassOutput {
	return &ClassPresenter{}
}

func (c ClassPresenter) HandleGetClasses(classes model.Classes, err error) (dto.GetClassesOutputForm, error) {
	output := dto.GetClassesOutputForm{}
	for _, class := range classes {
		form := &dto.GetClassByIdOutputForm{
			ID:        class.ID,
			LectureID: class.LectureID,
			Name:      class.Name,
			CreatedAt: class.CreatedAt,
			UpdatedAt: class.UpdatedAt,
		}
		output = append(output, form)
	}
	return output, err
}

func (c ClassPresenter) HandleGetClassById(class *model.Class, err error) (*dto.GetClassByIdOutputForm, error) {
	output := &dto.GetClassByIdOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}
	return output, err
}

func (c ClassPresenter) HandleCreateClass(class *model.Class, err error) (*dto.CreateClassOutputForm, error) {
	output := &dto.CreateClassOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}
	return output, err
}

func (c ClassPresenter) HandleUpdateClass(class *model.Class, err error) (*dto.UpdateClassOutputForm, error) {
	output := &dto.UpdateClassOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}
	return output, err
}

func (c ClassPresenter) HandleDeleteClass(count int64, err error) (*dto.DeleteClassOutputForm, error) {
	output := &dto.DeleteClassOutputForm{AffectedRowsCount: count}
	return output, err
}
