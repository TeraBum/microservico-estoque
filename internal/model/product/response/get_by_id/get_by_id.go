package getbyid

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type GetByIdResponse struct {
	Status      int
	Msg         string
	Id          uuid.UUID       `json:"id"`
	Name        string          `json:"name"`
	Description string          `json:"description"`
	Price       int64           `json:"price"`
	Category    string          `json:"category"`
	ImagesJson  json.RawMessage `json:"images_json"`
	IsActive    bool            `json:"is_active"`
	CreatedAt   time.Time       `json:"created_at"`
}
