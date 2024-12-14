package actions

import (
	"skeleton/repositories"
	"skeleton/requests"
	"skeleton/services"

	"github.com/google/uuid"
)

func (obj Test) Create(jwt services.JWTModel, request requests.TestCreate) (model repositories.TestModel) {
	model.Id = uuid.New()

	model.UserId = jwt.JWT.Id

	model.Name = request.Name

	model.Create()

	return
}
