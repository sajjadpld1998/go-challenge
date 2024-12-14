package requests

import (
	"github.com/gin-gonic/gin"
	"user_segmentation_service/helpers"
)

func (obj *UserSegmentPair) Validate(context *gin.Context) {
	request := helpers.RequestJsonBody(context)

	obj.UserId = request.StringValue("user_id")
	obj.Segment = request.StringValue("segment")

	ValidateRequestBody(obj)
}
