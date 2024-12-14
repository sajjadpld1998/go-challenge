package error_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func responseError(context *gin.Context, obj any, code int) {
	response := ResponseError{}

	response.Status = false
	response.Code = code

	responseError := ResponseErrorDetail{}
	responseError.Message = responseeMessage(code)
	responseError.Detail = obj

	response.Error = responseError

	context.JSON(code, response)
}

func responseeMessage(code int) (message string) {
	switch code {
	case http.StatusUnprocessableEntity:
		message = "Unprocessable Entity"
		break
	case http.StatusNotFound:
		message = "Not Found"
		break
	case http.StatusForbidden:
		message = "Accesss Forbidden"
		break
	case http.StatusUnauthorized:
		message = "Unauthorized"
		break
	case http.StatusInternalServerError:
		message = "Server Error"
		break
	case http.StatusNotAcceptable:
		message = "Not Acceptable"
		break
	}

	return
}

func responseUnprocessableEntity(context *gin.Context, obj any) {
	responseError(context, obj, http.StatusUnprocessableEntity)
}

func responseNotFound(context *gin.Context, obj any) {
	responseError(context, obj, http.StatusNotFound)
}

func responseForbidden(context *gin.Context, obj any) {
	responseError(context, obj, http.StatusForbidden)
}

func responseUnauthorized(context *gin.Context, obj any) {
	responseError(context, obj, http.StatusUnauthorized)
}

func responseServerError(context *gin.Context, obj any) {
	responseError(context, obj, http.StatusInternalServerError)
}
