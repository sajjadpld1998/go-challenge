package middleware

import (
	"skeleton/error_handler"
	"skeleton/services"

	"github.com/gin-gonic/gin"
)

func Authorization(context *gin.Context) {
	jwt := services.Jwt{}.Init(context)

	if !jwt.Check {
		context.Abort()

		error_handler.ThrowUnauthorizedError()
	}

	context.Next()
}

func AdminRole(context *gin.Context) {
	jwt := services.Jwt{}.Init(context)

	if !jwt.Check {
		context.Abort()

		error_handler.ThrowUnauthorizedError()
	}

	if jwt.JWT.Role != "admin" {
		context.Abort()

		error_handler.ThrowForbiddenAccessError()
	}

	context.Next()

}
