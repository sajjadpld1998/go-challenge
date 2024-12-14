package responses

import "github.com/google/uuid"

type Test struct {
	// building id. it's UUID id and is string
	Id uuid.UUID `json:"id" example:"a4d50e88-1a10-4a8b-af08-3e1060c00d5f"`
	// building name
	Name string `json:"name" example:"sci-fi1"`
	// creation time
	CreatedAt int64 `json:"created_at"`
}
