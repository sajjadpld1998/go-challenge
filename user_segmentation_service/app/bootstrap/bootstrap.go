package bootstrap

import (
	"user_segmentation_service/config"
	"user_segmentation_service/requests"
	"user_segmentation_service/routes"
)

func Boot() {
	//Load config file for using in application
	config.LoadConfig()

	//Initiate Gin instans and config this
	initGin()

	//Initiate request layer boot
	requests.InitValidation()

	//Initiate Routers
	routes.InitialRoutes(ginEngine)

	//Run project on a port and host
	runGin()
}
