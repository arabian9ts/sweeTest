// +build wireinject
package config

import (
	"github.com/google/wire"
)

var superSet = wire.NewSet(
	NewSettings,
	NewRDBMSSettings,
)

func provideSettings() (*Settings, error) {
	wire.Build(providerSet)
	return &Settings{}, nil
}
