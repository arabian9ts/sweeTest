package builder

//import (
//	"github.com/arabian9ts/sweeTest/app/interface/controllers"
//	"github.com/arabian9ts/sweeTest/app/interface/database"
//	"github.com/arabian9ts/sweeTest/app/interface/presenter"
//	"github.com/arabian9ts/sweeTest/app/usecase/interactor"
//	"github.com/arabian9ts/sweeTest/app/usecase/repository"
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
//var controllerSet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	database.NewLectureRepository,
//	controllers.NewStudentsController,
//	controllers.NewTasController,
//	controllers.NewTeachersController,
//	controllers.NewLecturesController,
//	controllers.NewRootController,
//	interactor.NewUserInteractor,
//	presenter.NewUserPresenter,
//	presenter.NewLecturePresenter,
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
//func InitializeStudentsController() (*controllers.StudentsController, error) {
//	wire.Build(controllerSet)
//	return &controllers.StudentsController{}, nil
//}
//
//func InitializeTasController() (*controllers.TasController, error) {
//	wire.Build(controllerSet)
//	return &controllers.TasController{}, nil
//}
//
//func InitializeTeachersController() (*controllers.TeachersController, error) {
//	wire.Build(controllerSet)
//	return &controllers.TeachersController{}, nil
//}
//
//func InitializeLecturesController() (*controllers.LecturesController, error) {
//	wire.Build(controllerSet)
//	return &controllers.LecturesController{}, nil
//}
