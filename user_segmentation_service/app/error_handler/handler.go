package error_handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func throwError(errorStruct interface{}) {
	panic(errorStruct)
}

func CatchError(context *gin.Context, errorStruct interface{}) {
	if errorStruct, ok := errorStruct.([]bodyValidationError); ok {
		responseUnprocessableEntity(context, errorStruct)
	}

	if errorStruct, ok := errorStruct.([]routeValidationError); ok {
		responseNotFound(context, errorStruct)
	}

	if errorStruct, ok := errorStruct.(forbiddenError); ok {
		arrayError := []forbiddenError{errorStruct}
		responseForbidden(context, arrayError)
	}

	if errorStruct, ok := errorStruct.(unauthorizedError); ok {
		arrayError := []unauthorizedError{errorStruct}
		responseUnauthorized(context, arrayError)
	}

	if errorStruct, ok := errorStruct.(serverError); ok {
		arrayError := []serverError{errorStruct}

		responseServerError(context, arrayError)
	}
}

func ThrowServerError(errMessage error) {
	server := serverError{
		Message: "Server under maintnance, we will be back soon.",
		err:     errMessage,
	}
	throwError(server)
}

func ThrowBodyValidationError(errors []validator.FieldError, messageGenerate messageGenerator) {
	var errorStructs []bodyValidationError

	for _, err := range errors {
		errorStruct := bodyValidationError{
			StructField: err.StructField(),
			ActualTag:   err.ActualTag(),
			Value:       err.Value(),
			Param:       err.Param(),
			Message:     messageGenerate(err.ActualTag(), err.StructField(), err.Param()),
		}

		errorStructs = append(errorStructs, errorStruct)
	}

	throwError(errorStructs)
}

func ThrowRouteValidationError(errors []validator.FieldError) {
	var errorStructs []routeValidationError

	for _, err := range errors {
		errorStruct := routeValidationError{
			StructField: err.StructField(),
			ActualTag:   err.ActualTag(),
			Message:     "The selected item does not exist!",
		}

		errorStructs = append(errorStructs, errorStruct)
	}

	throwError(errorStructs)
}

func ThrowForbiddenAccessError() {
	forbidden := forbiddenError{
		Message: "Access denied!",
	}

	throwError(forbidden)
}

func ThrowUnauthorizedError() {
	unauthorized := unauthorizedError{
		Message: "Unauthorized!",
	}

	throwError(unauthorized)
}
