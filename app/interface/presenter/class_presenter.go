package presenter

import (
	"github.com/arabian9ts/sweeTest/app/domain/model"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
)

type ClassPresenter struct{}

func NewClassPresenter() port.ClassOutput {
	return &ClassPresenter{}
}

func (*ClassPresenter) HandleGetClasses(classes model.Classes, err error) (dto.GetClassesOutputForm, error) {
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

func (*ClassPresenter) HandleGetClassById(class *model.Class, err error) (*dto.GetClassByIdOutputForm, error) {
	return &dto.GetClassByIdOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}, err
}

func (*ClassPresenter) HandleCreateClass(class *model.Class, err error) (*dto.CreateClassOutputForm, error) {
	return &dto.CreateClassOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}, err
}

func (*ClassPresenter) HandleUpdateClass(class *model.Class, err error) (*dto.UpdateClassOutputForm, error) {
	return &dto.UpdateClassOutputForm{
		ID:        class.ID,
		LectureID: class.LectureID,
		Name:      class.Name,
		CreatedAt: class.CreatedAt,
		UpdatedAt: class.UpdatedAt,
	}, err
}

func (*ClassPresenter) HandleDeleteClass(count int64, err error) (*dto.DeleteClassOutputForm, error) {
	return &dto.DeleteClassOutputForm{
		AffectedRowsCount: count,
	}, err
}
