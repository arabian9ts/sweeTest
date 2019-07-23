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
//)
//
//var controllerSet = wire.NewSet(
//	infrastructure.NewSqlHandler,
//	database.NewUserRepository,
//	controllers.NewStudentsController,
//	controllers.NewTasController,
//	controllers.NewTeachersController,
//	interactor.NewUserInteractor,
//	presenter.NewUserPresenter,
//)
//
//func InitializeUserRepository() (repository.UserRepository, error) {
//	wire.Build(repositorySet)
//	return &database.UserRepository{}, nil
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