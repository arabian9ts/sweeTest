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
//	controllers.NewUsersController,
//	interactor.NewUserInteractor,
//	presenter.NewUserPresenter,
//)
//
//func InitializeUserRepository() (repository.UserRepository, error) {
//	wire.Build(repositorySet)
//	return &database.UserRepository{}, nil
//}
//
//func InitializeUsersController() (*controllers.StudentsController, error) {
//	wire.Build(controllerSet)
//	return &controllers.StudentsController{}, nil
//}
