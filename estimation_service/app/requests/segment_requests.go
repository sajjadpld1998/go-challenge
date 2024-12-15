package requests

import (
	"estimation_service/helpers"
	"github.com/gin-gonic/gin"
)

func (obj *UserSegmentPair) Validate(context *gin.Context) {
	request := helpers.RequestJsonBody(context)

	obj.UserId = request.StringValue("user_id")
	obj.Segment = request.StringValue("segment")

	ValidateRequestBody(obj)
}

func (obj *SegmentRouteParam) Validate(context *gin.Context) {
	obj.Segment = context.Param("segment")

	ValidateRouteParam(obj)
}
