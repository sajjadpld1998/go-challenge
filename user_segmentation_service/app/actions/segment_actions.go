package actions

import (
	"user_segmentation_service/error_handler"
	"user_segmentation_service/requests"
	"user_segmentation_service/services"
)

func (obj Segment) ProcessUserSegmentPair(request requests.UserSegmentPair) {
	httpRequest := services.Http{}

	err := httpRequest.SendPairDataToES(request.UserId, request.Segment)

	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return
}
