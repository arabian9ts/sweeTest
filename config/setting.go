package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
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
	environment := os.Getenv("SWEETEST_ENV")
	if len(environment) == 0 {
		environment = "development"
	}

	// Note: if use binary in container, envPath is ./config/envs/{SWEETEST_ENV}.yml
	// otherwise, /{project absolute path to project}/config/envs/{SWEETEST_ENV}.yml
	project := "sweeTest"
	prefix := project + "/"
	configPath := fmt.Sprintf("config/envs/%s.yml", environment)
	envPath := "./" + configPath
	
	currentDir, err := os.Getwd()
	if err == nil && strings.Contains(currentDir, project) {
		rep := regexp.MustCompile(project + `.*`)
		envPath = rep.ReplaceAllString(currentDir, prefix + configPath)
	}

	buf, err := ioutil.ReadFile(envPath)
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
