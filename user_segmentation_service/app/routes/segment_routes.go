package routes

import (
	"github.com/gin-gonic/gin"
	"user_segmentation_service/controllers"
)

func (obj Segment) Routes(route *gin.Engine) {
	segmentController := controllers.Segment{}

	segmentGroup := route.Group("segments")
	{
		segmentGroup.POST("", segmentController.Create)
	}
}
