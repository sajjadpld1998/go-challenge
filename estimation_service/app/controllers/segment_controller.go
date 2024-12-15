package controllers

import (
	"estimation_service/actions"
	"estimation_service/requests"
	"estimation_service/responses"
	"github.com/gin-gonic/gin"
)

// Store godoc
// @Summary      Add a new segment pair
// @Description  Create new segment pair endpoint
// @Tags         Segment
// @Accept       json
// @Produce      json
// @Param        request   body      requests.UserSegmentPair  true "query params"
// @Success      201  {object}  responses.ResponseNoContent
// @Failure      422  {object}  error_handler.ResponseError{error=error_handler.ResponseErrorDetail{detail=[]error_handler.bodyValidationError}}
// @Failure      500  {object}  error_handler.ResponseError{error=error_handler.ResponseErrorDetail{detail=[]error_handler.serverError}}
// @Router       /segments [post]
func (obj Segment) Store(context *gin.Context) {
	request := requests.UserSegmentPair{}
	request.Validate(context)

	action := actions.Segment{}
	action.StoreUserSegmentPair(request)

	responses.ResponseSuccess(context)
}

// Show godoc
// @Summary      Show a Segment count
// @Description  Show a Segment count by it's name
// @Tags         Segment
// @Accept       json
// @Produce      json
// @Param        segment   path      string  true  "segment"
// @Success      200  {object}  responses.ResponseContent{data=responses.SegmentUsersCount}
// @Failure      404  {object}  error_handler.ResponseError{error=error_handler.ResponseErrorDetail{detail=[]error_handler.routeValidationError}}
// @Failure      500  {object}  error_handler.ResponseError{error=error_handler.ResponseErrorDetail{detail=[]error_handler.serverError}}
// @Router       /segments/{segment} [get]
func (obj Segment) Show(context *gin.Context) {
	request := requests.SegmentRouteParam{}
	request.Validate(context)

	action := actions.Segment{}
	count := action.GetSegmentUserCount(request)

	response := responses.SegmentUsersCount{}
	response.Response(context, count)
}
