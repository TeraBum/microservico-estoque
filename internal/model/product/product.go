package product

import (
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	Id          uuid.UUID       `db:"Id" json:"id"`
	Name        string          `db:"Name" json:"name"`
	Description string          `db:"Description" json:"description"`
	Price       int64           `db:"Price" json:"price"`
	Category    string          `db:"Category" json:"category"`
	ImagesJson  json.RawMessage `db:"ImagesJson" json:"images_json"`
	IsActive    bool            `db:"IsActive" json:"is_active"`
	CreatedAt   *time.Time      `db:"CreatedAt" json:"created_at,omitempty"`
}
