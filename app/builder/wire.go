package builder

//import (
//	"github.com/arabian9ts/sweeTest/app/interface/controllers"
//	"github.com/arabian9ts/sweeTest/app/interface/database"
//	"github.com/arabian9ts/sweeTest/app/interface/handler"
//	"github.com/arabian9ts/sweeTest/app/interface/presenter"
//	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
//	"github.com/arabian9ts/sweeTest/app/usecase/repository"
//	"github.com/arabian9ts/sweeTest/app/validator"
//	"github.com/arabian9ts/sweeTest/infrastructure"
//	"github.com/google/wire"
//)
//
//var repositorySet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	database.NewLectureRepository,
//)
//
//var handlerSet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	handler.NewAuthHandler,
//	handler.NewMeHandler,
//	handler.NewRootHandler,
//)
//
//var controllerSet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	database.NewLectureRepository,
//	database.NewTaskRepository,
//	controllers.NewStudentsController,
//	controllers.NewAssistantsController,
//	controllers.NewTeachersController,
//	controllers.NewLecturesController,
//	controllers.NewRootController,
//	controllers.NewTasksController,
//	controllers.NewStudentLoginController,
//	controllers.NewAssistantLoginController,
//	controllers.NewTeacherLoginController,
//	controllers.NewAdminLoginController,
//	interactor.NewUserInteractor,
//	interactor.NewLectureInteractor,
//	interactor.NewTaskInteractor,
//	interactor.NewAuthInteractor,
//	presenter.NewUserPresenter,
//	presenter.NewLecturePresenter,
//	presenter.NewTaskPresenter,
//	presenter.NewAuthPresenter,
//	validator.NewDefaultValidator,
//)
//
//func InitializeUserRepository() (repository.UserRepository, error) {
//	wire.Build(repositorySet)
//	return &database.UserRepository{}, nil
//}
//
//func InitializeLectureRepository() (repository.LectureRepository, error) {
//	wire.Build(repositorySet)
//	return &database.LectureRepository{}, nil
//}
//
//func InitializeRootController() (*controllers.RootController, error) {
//	wire.Build(controllerSet)
//	return &controllers.RootController{}, nil
//}
//
//func InitializeRootHandler() (*handler.RootHandler, error) {
//	wire.Build(handlerSet)
//	return &handler.RootHandler{}, nil
//}
