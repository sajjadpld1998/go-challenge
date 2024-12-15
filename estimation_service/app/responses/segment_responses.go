package responses

import (
	"github.com/gin-gonic/gin"
)

func (obj *SegmentUsersCount) Response(context *gin.Context, count int) {
	obj.Count = count

	ResponseSuccessWithContent(context, obj)
}
