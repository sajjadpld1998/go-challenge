package routes

import (
	"estimation_service/controllers"
	"github.com/gin-gonic/gin"
)

func (obj Segment) Routes(route *gin.Engine) {
	segmentController := controllers.Segment{}

	segmentGroup := route.Group("segments")
	{
		segmentGroup.POST("", segmentController.Store)
		segmentGroup.GET(":segment", segmentController.Show)
	}
}
