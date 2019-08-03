// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package builder

import (
	"github.com/arabian9ts/sweeTest/app/interface/controllers"
	"github.com/arabian9ts/sweeTest/app/interface/database"
	"github.com/arabian9ts/sweeTest/app/interface/handler"
	"github.com/arabian9ts/sweeTest/app/interface/presenter"
	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
	"github.com/arabian9ts/sweeTest/app/usecase/repository"
	"github.com/arabian9ts/sweeTest/app/validator"
	"github.com/arabian9ts/sweeTest/infrastructure"
	"github.com/google/wire"
)

// Injectors from wire.go:

func InitializeUserRepository() (repository.UserRepository, error) {
	sqlHandler := infrastructure.NewSqlHandler()
	userRepository, err := database.NewUserRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	return userRepository, nil
}

func InitializeLectureRepository() (repository.LectureRepository, error) {
	sqlHandler := infrastructure.NewSqlHandler()
	lectureRepository, err := database.NewLectureRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	return lectureRepository, nil
}

func InitializeRootController() (*controllers.RootController, error) {
	sqlHandler := infrastructure.NewSqlHandler()
	userRepository, err := database.NewUserRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	userOutput := presenter.NewUserPresenter()
	validation := validator.NewDefaultValidator()
	studentsController, err := controllers.NewStudentsController(userRepository, userOutput, validation)
	if err != nil {
		return nil, err
	}
	assistantsController, err := controllers.NewAssistantsController(userRepository, userOutput, validation)
	if err != nil {
		return nil, err
	}
	teachersController, err := controllers.NewTeachersController(userRepository, userOutput, validation)
	if err != nil {
		return nil, err
	}
	lectureRepository, err := database.NewLectureRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	lectureOutput := presenter.NewLecturePresenter()
	lecturesController, err := controllers.NewLecturesController(lectureRepository, lectureOutput, validation)
	if err != nil {
		return nil, err
	}
	taskRepository, err := database.NewTaskRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	taskOutput := presenter.NewTaskPresenter()
	tasksController, err := controllers.NewTasksController(taskRepository, taskOutput, validation)
	if err != nil {
		return nil, err
	}
	authOutput := presenter.NewAuthPresenter()
	studentLoginController, err := controllers.NewStudentLoginController(userRepository, authOutput, validation)
	if err != nil {
		return nil, err
	}
	assistantLoginController, err := controllers.NewAssistantLoginController(userRepository, authOutput, validation)
	if err != nil {
		return nil, err
	}
	teacherLoginController, err := controllers.NewTeacherLoginController(userRepository, authOutput, validation)
	if err != nil {
		return nil, err
	}
	adminLoginController, err := controllers.NewAdminLoginController(userRepository, authOutput, validation)
	if err != nil {
		return nil, err
	}
	helpOutput := presenter.NewHelpPresenter()
	helpRepository, err := database.NewHelpRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	helpsController, err := controllers.NewHelpsController(helpOutput, helpRepository, validation)
	if err != nil {
		return nil, err
	}
	commentOutput := presenter.NewCommentPresenter()
	commentRepository := database.NewCommentRepository(sqlHandler)
	studentCommentsController, err := controllers.NewStudentCommentsController(commentOutput, commentRepository, validation)
	if err != nil {
		return nil, err
	}
	rootController, err := controllers.NewRootController(studentsController, assistantsController, teachersController, lecturesController, tasksController, studentLoginController, assistantLoginController, teacherLoginController, adminLoginController, helpsController, studentCommentsController)
	if err != nil {
		return nil, err
	}
	return rootController, nil
}

func InitializeRootHandler() (*handler.RootHandler, error) {
	sqlHandler := infrastructure.NewSqlHandler()
	userRepository, err := database.NewUserRepository(sqlHandler)
	if err != nil {
		return nil, err
	}
	authHandler, err := handler.NewAuthHandler(userRepository)
	if err != nil {
		return nil, err
	}
	meHandler, err := handler.NewMeHandler(userRepository)
	if err != nil {
		return nil, err
	}
	rootHandler, err := handler.NewRootHandler(authHandler, meHandler)
	if err != nil {
		return nil, err
	}
	return rootHandler, nil
}

// wire.go:

var repositorySet = wire.NewSet(infrastructure.NewSqlHandler, database.NewUserRepository, database.NewLectureRepository)

var handlerSet = wire.NewSet(infrastructure.NewSqlHandler, database.NewUserRepository, handler.NewAuthHandler, handler.NewMeHandler, handler.NewRootHandler)

var controllerSet = wire.NewSet(infrastructure.NewSqlHandler, database.NewUserRepository, database.NewLectureRepository, database.NewTaskRepository, database.NewHelpRepository, database.NewCommentRepository, controllers.NewStudentsController, controllers.NewAssistantsController, controllers.NewTeachersController, controllers.NewLecturesController, controllers.NewRootController, controllers.NewTasksController, controllers.NewStudentLoginController, controllers.NewAssistantLoginController, controllers.NewTeacherLoginController, controllers.NewAdminLoginController, controllers.NewHelpsController, controllers.NewStudentCommentsController, controllers.NewAssistantCommentsController, controllers.NewTeacherCommentsController, interactor.NewUserInteractor, interactor.NewLectureInteractor, interactor.NewTaskInteractor, interactor.NewAuthInteractor, interactor.NewHelpInteractor, interactor.NewCommentInteractor, presenter.NewUserPresenter, presenter.NewLecturePresenter, presenter.NewTaskPresenter, presenter.NewAuthPresenter, presenter.NewHelpPresenter, presenter.NewCommentPresenter, validator.NewDefaultValidator)
