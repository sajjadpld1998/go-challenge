package actions

import (
	"user_segmentation_service/repositories"
	"user_segmentation_service/requests"
	"user_segmentation_service/services"

	"github.com/google/uuid"
)

func (obj Test) Create(jwt services.JWTModel, request requests.TestCreate) (model repositories.TestModel) {
	model.Id = uuid.New()

	model.UserId = jwt.JWT.Id

	model.Name = request.Name

	model.Create()

	return
}
