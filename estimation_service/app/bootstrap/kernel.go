package bootstrap

import (
	"github.com/getsentry/sentry-go"
	"skeleton/config"
	"skeleton/error_handler"

	"github.com/gin-gonic/gin"
)

var ginEngine *gin.Engine

func initGin() {
	//Set gin engin mode (debug-release-test)
	gin.SetMode(config.GetConfig().AppENV)

	//Make new isntance of Gin and share it in the project
	ginEngine = gin.Default()

	//Maximum request size
	ginEngine.MaxMultipartMemory = 100 << 20 // 100 MB

	//Patch Error handling core to Gin for responsing errors to client
	ginEngine.Use(gin.CustomRecovery(error_handler.CatchError))
}

func runGin() {
	//Run Gin on especial port and host
	ginEngine.Run(config.GetConfig().Host)
}
func InitSentry() {
	sentry.Init(sentry.ClientOptions{
		Dsn:         config.GetConfig().Sentry.Dsn,
		Environment: config.GetConfig().AppENV,
	})
}