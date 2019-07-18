package config

import (
	"fmt"
	"github.com/go-yaml/yaml"
	"io/ioutil"
	"os"
)

type RDBMS struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	DataBase string `yaml:"database"`
	Dialect  string `yaml:"dialect"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func NewRDBMSSettings() (*RDBMS, error) {
	environment := os.Getenv("GO_ENV")
	if len(environment) == 0 {
		environment = "development"
	}

	buf, err := ioutil.ReadFile(fmt.Sprintf("config/envs/%s.yml", environment))
	if err != nil {
		panic("failed to load env settings")
	}

	rdb := &RDBMS{}
	err = yaml.Unmarshal(buf, rdb)
	if err != nil {
		panic("failed to map settings to struct")
	}

	return rdb, nil
}

