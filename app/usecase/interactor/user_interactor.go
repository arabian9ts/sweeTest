package interactor

import (
	"github.com/arabian9ts/sweeTest/app/adapter"
	"github.com/arabian9ts/sweeTest/app/dto"
	"github.com/arabian9ts/sweeTest/app/usecase/port"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
)

type UserInteractor struct {
	UserRepository repository.UserRepository
	UserOutput     port.UserOutput
}

func NewUserInteractor(repository repository.UserRepository, output port.UserOutput) (*UserInteractor, error) {
	return &UserInteractor{UserRepository: repository, UserOutput: output}, nil
}

func (interactor *UserInteractor) GetStudents(limit int, offset int) (dto.GetStudentsOutputForm, error) {
	return interactor.UserOutput.HandleGetStudents(
		interactor.UserRepository.GetStudents(limit, offset),
	)
}

func (interactor *UserInteractor) GetAssistants(limit int, offset int) (dto.GetAssistantsOutputForm, error) {
	return interactor.UserOutput.HandleGetAssistants(
		interactor.UserRepository.GetAssistants(limit, offset),
	)
}

func (interactor *UserInteractor) GetTeachers(limit int, offset int) (dto.GetTeachersOutputForm, error) {
	return interactor.UserOutput.HandleGetTeachers(
		interactor.UserRepository.GetTeachers(limit, offset),
	)
}

func (interactor *UserInteractor) GetStudentById(id int64) (*dto.GetStudentOutputForm, error) {
	return interactor.UserOutput.HandleGetStudent(
		interactor.UserRepository.GetStudentById(id),
	)
}

func (interactor *UserInteractor) GetAssistantById(id int64) (*dto.GetAssistantOutputForm, error) {
	return interactor.UserOutput.HandleGetAssistant(
		interactor.UserRepository.GetAssistantById(id),
	)
}

func (interactor *UserInteractor) GetTeacherById(id int64) (*dto.GetTeacherOutputForm, error) {
	return interactor.UserOutput.HandleGetTeacher(
		interactor.UserRepository.GetTeacherById(id),
	)
}

func (interactor *UserInteractor) GetParticipatedStudents(lectureID int64, limit int, offset int) (dto.GetStudentsOutputForm, error) {
	return interactor.UserOutput.HandleGetStudents(
		interactor.UserRepository.GetParticipatingStudentsOfLecture(lectureID, limit, offset),
	)
}

func (interactor *UserInteractor) GetParticipatedAssistants(lectureID int64, limit int, offset int) (dto.GetAssistantsOutputForm, error) {
	return interactor.UserOutput.HandleGetAssistants(
		interactor.UserRepository.GetParticipatingAssistantsOfLecture(lectureID, limit, offset),
	)
}

func (interactor *UserInteractor) GetParticipatedTeachers(lectureID int64, limit int, offset int) (dto.GetTeachersOutputForm, error) {
	return interactor.UserOutput.HandleGetTeachers(
		interactor.UserRepository.GetParticipatingTeachersOfLecture(lectureID, limit, offset),
	)
}

func (interactor *UserInteractor) CreateStudent(form *dto.CreateStudentInputForm) (*dto.CreateStudentOutputForm, error) {
	student := adapter.ConvertCreateStudentInputFormToUser(form)
	return interactor.UserOutput.HandleCreateStudent(
		interactor.UserRepository.InsertStudent(student),
	)
}

func (interactor *UserInteractor) CreateAssistant(form *dto.CreateAssistantInputForm) (*dto.CreateAssistantOutputForm, error) {
	assistant := adapter.ConvertCreateAssistantInputFormToAssistant(form)
	return interactor.UserOutput.HandleCreateAssistant(
		interactor.UserRepository.InsertAssistant(assistant),
	)
}

func (interactor *UserInteractor) CreateTeacher(form *dto.CreateTeacherInputForm) (*dto.CreateTeacherOutputForm, error) {
	teacher := adapter.ConvertCreateTeacherInputFormToTeacher(form)
	return interactor.UserOutput.HandleCreateTeacher(
		interactor.UserRepository.InsertTeacher(teacher),
	)
}

func (interactor *UserInteractor) UpdateStudent(form *dto.UpdateStudentInputForm) (*dto.UpdateStudentOutputForm, error) {
	student := adapter.ConvertUpdateStudentInputFormToUser(form)
	return interactor.UserOutput.HandleUpdateStudent(
		interactor.UserRepository.UpdateStudent(student),
	)
}

func (interactor *UserInteractor) UpdateAssistant(form *dto.UpdateAssistantInputForm) (*dto.UpdateAssistantOutputForm, error) {
	assistant := adapter.ConvertUpdateAssistantInputFormToAssistant(form)
	return interactor.UserOutput.HandleUpdateAssistant(
		interactor.UserRepository.UpdateAssistant(assistant),
	)
}

func (interactor *UserInteractor) UpdateTeacher(form *dto.UpdateTeacherInputForm) (*dto.UpdateTeacherOutputForm, error) {
	teacher := adapter.ConvertUpdateTeacherInputFormToTeacher(form)
	return interactor.UserOutput.HandleUpdateTeacher(
		interactor.UserRepository.UpdateTeacher(teacher),
	)
}
