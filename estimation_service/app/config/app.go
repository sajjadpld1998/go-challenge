package config

import (
	"estimation_service/error_handler"

	"github.com/jinzhu/configor"
)

var loadedConfig *config

type config struct {
	AppName  string `default:"es"`
	AppENV   string `default:"debug"`
	Port     string `default:"8484"`
	Host     string `default:":8484"`
	Database struct {
		Redis struct {
			Password string `default:""`
			Host     string `default:"127.0.0.1:6379"`
		}
	}
}

func LoadConfig() {
	loadedConfig = &config{}
	err := configor.Load(loadedConfig, "configs.yaml")
	if err != nil {
		error_handler.ThrowServerError(err)
	}
}

func GetConfig() *config {
	return loadedConfig
}
