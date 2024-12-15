package routes

import (
	_ "estimation_service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Estimation Service
// @version			1.0
// @description	This service containes Estimation Service endpoints.
//
// @contact.name	Sajjad Pouladvand
// @contact.email	sajjadpld1998@gmail.com
//
// @host			localhost:8282
// @BasePath		/
func InitialRoutes(router *gin.Engine) {
	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	segment := Segment{}
	segment.Routes(router)
}
