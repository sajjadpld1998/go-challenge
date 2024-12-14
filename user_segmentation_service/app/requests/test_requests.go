package requests

import "github.com/gin-gonic/gin"

func (obj *TestCreate) Validate(context *gin.Context) {
	obj.Name = context.PostForm("name")

	ValidateRequestBody(obj)
}
