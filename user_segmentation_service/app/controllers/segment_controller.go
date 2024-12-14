package controllers

import (
	"github.com/gin-gonic/gin"
	"user_segmentation_service/actions"
	"user_segmentation_service/requests"
	"user_segmentation_service/responses"
)

// Create godoc
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
func (obj Segment) Create(context *gin.Context) {
	request := requests.UserSegmentPair{}
	request.Validate(context)

	action := actions.Segment{}
	action.ProcessUserSegmentPair(request)

	responses.ResponseSuccess(context)
}
