package actions

import (
	"estimation_service/error_handler"
	"estimation_service/repositories"
	"estimation_service/requests"
)

func (obj Segment) StoreUserSegmentPair(request requests.UserSegmentPair) {
	model := repositories.UserSegmentPair{}

	model.UserId = request.UserId
	model.Segment = request.Segment

	err := model.Store()
	if err != nil {
		error_handler.ThrowServerError(err)
	}
}

func (obj Segment) GetSegmentUserCount(request requests.SegmentRouteParam) (count int) {
	model := repositories.UserSegmentPair{}

	count, err := model.GetSegmentCount(request.Segment)
	if err != nil {
		error_handler.ThrowServerError(err)
	}

	return
}
