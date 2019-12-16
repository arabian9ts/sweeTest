package builder

//import (
//	"github.com/google/wire"
//
//	"github.com/arabian9ts/sweeTest/app/interface/controllers"
//	"github.com/arabian9ts/sweeTest/app/interface/database"
//	"github.com/arabian9ts/sweeTest/app/interface/handler"
//	"github.com/arabian9ts/sweeTest/app/interface/presenter"
//	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
//	"github.com/arabian9ts/sweeTest/app/usecase/repository"
//	"github.com/arabian9ts/sweeTest/app/validator"
//	"github.com/arabian9ts/sweeTest/infrastructure"
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
//	handler.NewAliveHandler,
//)
//
//var controllerSet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	database.NewLectureRepository,
//	database.NewTaskRepository,
//	database.NewHelpRepository,
//	database.NewCommentRepository,
//	database.NewParticipationRepository,
//	controllers.NewStudentsController,
//	controllers.NewAssistantsController,
//	controllers.NewTeachersController,
//	controllers.NewLecturesController,
//	controllers.NewRootController,
//	controllers.NewTasksController,
//	controllers.NewLoginController,
//	controllers.NewHelpsController,
//	controllers.NewCommentsController,
//	controllers.NewParticipationController,
//	controllers.NewParticipatingLecturesController,
//	controllers.NewParticipantsController,
//	interactor.NewUserInteractor,
//	interactor.NewLectureInteractor,
//	interactor.NewTaskInteractor,
//	interactor.NewAuthInteractor,
//	interactor.NewHelpInteractor,
//	interactor.NewCommentInteractor,
//	interactor.NewParticipationInteractor,
//	presenter.NewUserPresenter,
//	presenter.NewLecturePresenter,
//	presenter.NewTaskPresenter,
//	presenter.NewAuthPresenter,
//	presenter.NewHelpPresenter,
//	presenter.NewCommentPresenter,
//	presenter.NewParticipationPresenter,
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
