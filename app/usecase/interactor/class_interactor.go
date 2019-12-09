package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type ClassInteractor struct {
	ClassRepository repository.ClassRepository
	ClassOutput     port.ClassOutput
}

func NewClassInteractor(repository repository.ClassRepository, output port.ClassOutput) (*ClassInteractor, error) {
	return &ClassInteractor{ClassRepository: repository, ClassOutput: output}, nil
}

func (interactor *ClassInteractor) GetClassById(id int64) (*dto.GetClassByIdOutputForm, error) {
	return interactor.ClassOutput.HandleGetClassById(
		interactor.ClassRepository.GetClassById(id),
	)
}

func (interactor *ClassInteractor) GetClassesByLectureId(lectureID int64, limit int, offset int) (dto.GetClassesOutputForm, error) {
	return interactor.ClassOutput.HandleGetClasses(
		interactor.ClassRepository.GetClassesByLectureID(lectureID, limit, offset),
	)
}
func (interactor *ClassInteractor) CreateClass(form *dto.CreateClassInputForm) (*dto.CreateClassOutputForm, error) {
	class := adapter.ConvertCreateClassInputFormToClass(form)
	return interactor.ClassOutput.HandleCreateClass(
		interactor.ClassRepository.CreateClass(class),
	)
}

func (interactor *ClassInteractor) UpdateClass(form *dto.UpdateClassInputForm) (*dto.UpdateClassOutputForm, error) {
	class := adapter.ConvertUpdateClassInputFormToClass(form)
	return interactor.ClassOutput.HandleUpdateClass(
		interactor.ClassRepository.UpdateClass(class),
	)
}

func (interactor *ClassInteractor) DeleteClass(id int64, lectureID int64) (*dto.DeleteClassOutputForm, error) {
	return interactor.ClassOutput.HandleDeleteClass(
		interactor.ClassRepository.DeleteClass(id, lectureID),
	)
}
