package config

import (
	"skeleton/error_handler"

	"github.com/jinzhu/configor"
)

var loadedConfig *config

type config struct {
	AppName string `default:"landrocker"`
	AppENV  string `default:"debug"`
	Port    string `default:"8282"`
	Host    string `default:":8282"`
	Jwt     struct {
		Key string `default:"key"`
	}
	Sentry struct {
		Dsn string `default:"https://169825f408004d7eb77d8abf245684d4@o4505266876383232.ingest.sentry.io/4505266879594496"`
	}
	Path struct {
		ProductStorePathDirectory        string `default:"Product"`
		ProductImagesStorePathDirectory  string `default:"images"`
		KitBlockStorePathDirectory       string `default:"kit-bash"`
		KitBlockFileStorePathDirectory   string `default:"file"`
		KitBlockImagesStorePathDirectory string `default:"images"`
	}
	Database struct {
		LandrockerRead struct {
			Driver   string `default:"postgres"`
			User     string `default:"user7"`
			Password string `default:"pass7"`
			Host     string `default:"127.0.0.1"`
			Port     int    `default:"5432"`
			Ssl      string `default:"disable"`
			Database string `default:"test7"`
		}
		LandrockerWrite struct {
			Driver   string `default:"postgres"`
			User     string `default:"user7"`
			Password string `default:"pass7"`
			Host     string `default:"127.0.0.1"`
			Port     int    `default:"5432"`
			Ssl      string `default:"disable"`
			Database string `default:"test7"`
		}
	}
	Aws struct {
		Region          string `default:"aws_region"`
		AccessKeyId     string `default:"aws_key"`
		SecretAccessKey string `default:"aws_secret"`
	}
	FileSystem struct {
		Default string `default:"s3"`
		Disks   struct {
			Local struct {
				Driver string `default:"local"`
				Root   string `default:"/var/www/html/kitblock-service.landrocker.io/builder"`
				Url    string `default:"https://kitblock-service.landrocker.io/builder"`
			}
			Public struct {
				Driver string `default:"local"`
				Root   string `default:"/public"`
				Url    string `default:"https://test.io/public"`
			}
			S3 struct {
				Driver   string `default:"s3"`
				Endpoint string `default:"aws_endpoint"`
				Bucket   string `default:"/public"`
				Url      string `default:"https://test.io/public"`
			}
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
