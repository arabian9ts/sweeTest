// +build wireinject
package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"time"
)

var settings *Settings

func init() {
	var err error = nil
	settings, err = NewSettings()
	if err != nil {
		panic("settings not created correctly")
	}
}

type Settings struct {
	MySQL *MySQL
	Jwt   *Jwt
}

type MySQL struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type Jwt struct {
	Secret       string `yaml:"secret"`
	ValidMinutes int64  `yaml:"valid_minutes"`
	Realm        string `yaml:"realm"`
}

func (jwt *Jwt) ExpiresTime() time.Time {
	minutes := time.Minute * time.Duration(jwt.ValidMinutes)
	return time.Now().Add(minutes)
}

func NewSettings() (*Settings, error) {
	environment := os.Getenv("GOENV")
	if len(environment) == 0 {
		environment = "development"
	}

	buf, err := ioutil.ReadFile(fmt.Sprintf("config/envs/%s.yml", environment))
	if err != nil {
		panic("failed to load env settings")
	}

	settings := &Settings{}
	err = yaml.Unmarshal(buf, settings)
	if err != nil {
		fmt.Println(err)
		panic("failed to map settings to struct")
	}
	return settings, nil
}

func GetRdbUri() (uri string) {
	rdb := settings.MySQL
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

func GetJwtConfig() *Jwt {
	return settings.Jwt
}
