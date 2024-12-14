package repositories

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TestModel struct {
	// it's UUID id and is string
	Id uuid.UUID `gorm:"primaryKey;type:uuid;uniqueIndex:building_id_index;not null"`
	// user id. it's UUID id and is string
	UserId uuid.UUID `gorm:"type:uuid;index:building_user_id_index;"`
	// name
	Name string `gorm:"type:varchar(255);index:building_name_index;not null"`
	// create date time
	CreatedAt time.Time `gorm:"autoCreateTime"`
	// update date time
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
	// date of deleting
	DeletedAt gorm.DeletedAt `gorm:"softDelete"`
}
