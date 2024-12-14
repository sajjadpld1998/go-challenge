package repositories

import "skeleton/error_handler"

func (obj *TestModel) Create() {
	result := WriteConnection.Create(&obj)

	if result.Error != nil {
		error_handler.ThrowServerError(result.Error)
	}
}

func (obj TestModel) Update(id string, updateItems map[string]interface{}) {
	result := WriteConnection.Model(&TestModel{}).Where("id = ? ", id).Updates(updateItems)

	if result.Error != nil {
		error_handler.ThrowServerError(result.Error)
	}
}