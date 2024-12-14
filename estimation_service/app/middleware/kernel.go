package middleware

import (
	"github.com/gin-gonic/gin"
)

var routeMiddlewares = make(map[string]gin.HandlerFunc)

func initGlobalMiddlewares(ginEngine *gin.Engine) {
	//add global middlewares here
}

func InitMiddleware(ginEngine *gin.Engine) {
	initGlobalMiddlewares(ginEngine)
}
