package responses

import (
	"skeleton/repositories"

	"github.com/gin-gonic/gin"
)

func (obj *Test) mapRepositoryToResponse(model repositories.TestModel) {
	obj.Id = model.Id
	obj.Name = model.Name
}

func (obj *Test) ResponseCreateBuilding(context *gin.Context, model repositories.TestModel) {
	obj.mapRepositoryToResponse(model)

	ResponseCreatedWithContent(context, obj)
}
