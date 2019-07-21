// +build wireinject
package config

import "fmt"

var settings *Settings

func init() {
	var err error = nil
	settings, err = initializeSettings()
	if err != nil {
		panic("settings not created correctly")
	}
}

type Settings struct {
	*RDBMS
}

func NewSettings(rdb *RDBMS) (*Settings, error) {
	return &Settings{ RDBMS: rdb }, nil
}

func GetSettings() *Settings {
	return settings
}

func (setting *Settings) GetRdbUri() (uri string) {
	rdb := settings.RDBMS
	uri = fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		rdb.UserName,
		rdb.Password,
		rdb.Host,
		rdb.Port,
		rdb.DataBase,
	)
	return
}
