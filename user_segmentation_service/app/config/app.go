package config

import (
	"user_segmentation_service/error_handler"

	"github.com/jinzhu/configor"
)

var loadedConfig *config

type config struct {
	AppName string `default:"uss"`
	AppENV  string `default:"debug"`
	Port    string `default:"8282"`
	Host    string `default:":8282"`
	Url     struct {
		EstimateServiceStorePair string `default:"http://localhost:8484/segments"`
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
