package builder

import (
	"github.com/arabian9ts/sweeTest/app/repository"
	"github.com/arabian9ts/sweeTest/app/interface/database"
	"github.com/arabian9ts/sweeTest/infrastructure"
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	infrastructure.NewSqlHandler,
	database.NewUserRepository,
)

func ProvideUserRepository() (repository.UserRepository, error) {
	wire.Build(superSet)
	return &database.UserRepository{}, nil
}
