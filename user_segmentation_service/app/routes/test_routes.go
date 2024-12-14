package routes

import (
	"user_segmentation_service/controllers"
	"user_segmentation_service/middleware"

	"github.com/gin-gonic/gin"
)

func (obj Test) Routes(route *gin.Engine) {
	testController := controllers.Test{}

	buildingGroup := route.Group("test")
	{
		buildingGroup.Use(middleware.Authorization).POST("", testController.Create)
	}
}
