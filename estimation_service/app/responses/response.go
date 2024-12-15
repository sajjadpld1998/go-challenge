package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ResponseSuccess(context *gin.Context) {
	response := ResponseNoContent{}

	response.Status = true
	response.Code = http.StatusOK

	context.JSON(http.StatusOK, response)
}

func ResponseSuccessWithContent(context *gin.Context, obj any) {
	response := ResponseContent{}

	response.Status = true
	response.Code = http.StatusOK
	response.Data = obj

	context.JSON(http.StatusOK, response)
}

func ResponseCreated(context *gin.Context) {
	response := ResponseNoContent{}

	response.Status = true
	response.Code = http.StatusCreated

	context.JSON(http.StatusCreated, response)
}

func ResponseCreatedWithContent(context *gin.Context, obj any) {
	response := ResponseContent{}

	response.Status = true
	response.Code = http.StatusCreated
	response.Data = obj

	context.JSON(http.StatusCreated, response)
}

func ResponseUpdated(context *gin.Context) {
	response := ResponseNoContent{}

	response.Status = true
	response.Code = http.StatusOK

	context.JSON(http.StatusOK, response)
}

func ResponseDeleted(context *gin.Context) {
	response := ResponseNoContent{}

	response.Status = true
	response.Code = http.StatusOK

	context.JSON(http.StatusOK, response)
}
