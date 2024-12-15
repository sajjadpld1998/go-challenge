package bootstrap

import (
	"estimation_service/config"
	"estimation_service/repositories"
	"estimation_service/requests"
	"estimation_service/routes"
	"estimation_service/services"
)

func Boot() {
	//Load config file for using in application
	config.LoadConfig()

	//Initiate Gin instans and config this
	initGin()

	//Initiate services
	services.InitServices()

	//Initiate connection to database
	repositories.InitDBConnection()

	//Initiate request layer boot
	requests.InitValidation()

	//Initiate Routers
	routes.InitialRoutes(ginEngine)

	//Run project on a port and host
	runGin()
}
