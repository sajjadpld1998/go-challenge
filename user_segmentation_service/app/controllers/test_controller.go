package controllers

import (
	"skeleton/actions"
	"skeleton/requests"
	"skeleton/responses"
	"skeleton/services"

	"github.com/gin-gonic/gin"
)

// Create godoc
// @Summary      Add a new Test
// @Description  Create new Test endpoint
// @Tags         Test
// @Accept       mpfd
// @Produce      json
// @Param        Authorization   header      string  true  "jwt token" example(eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZ.........)
// @Param        name   formData      string  true  "name of building. it must be unique" example(home_building)
// @Success      201  {object}  responses.ResponseContent{data=building_responses.BuildingResponce}
// @Failure      401  {object}  error_handler.ResponseContent{data=error_handler.ResponseStaticError{detail=[]error_handler.unauthorizedError}}
// @Failure      422  {object}  error_handler.ResponseContent{data=error_handler.ResponseStaticError{detail=[]error_handler.bodyValidationError}}
// @Failure      500  {object}  error_handler.ResponseContent{data=error_handler.ResponseStaticError{detail=[]error_handler.serverError}}
// @Router       /building [post]
func (obj Test) Create(context *gin.Context) {
	request := requests.TestCreate{}
	request.Validate(context)

	jwt := services.Jwt{}.Init(context)

	action := actions.Test{}
	model := action.Create(jwt, request)

	response := responses.Test{}

	response.ResponseCreateBuilding(context, model)
}
