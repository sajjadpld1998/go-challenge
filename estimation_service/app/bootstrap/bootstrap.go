package bootstrap

import (
	"skeleton/config"
	"skeleton/middleware"
	"skeleton/repositories"
	"skeleton/requests"
	"skeleton/routes"
	"skeleton/services"
)

func Boot() {
	//Load config file for using in application
	config.LoadConfig()

	//Initiate Gin instans and config this
	initGin()

	//init sentry
	InitSentry()

	//Initiate services
	services.InitServices()

	//Initiate connection to database
	repositories.InitDBConnection()

	//Initiate request layer boot
	requests.InitValidation()

	//Initiate middlewares
	middleware.InitMiddleware(ginEngine)

	//Initiate Routers
	routes.InitialRoutes(ginEngine)

	//Run project on a port and host
	runGin()
}
