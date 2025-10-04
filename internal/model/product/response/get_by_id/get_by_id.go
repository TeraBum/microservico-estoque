package getbyid

import (
	"time"

	"github.com/gofrs/uuid"
)

type GetByIdResponse struct {
	Status      int        `json:"-"`
	Msg         string     `json:"-"`
	Id          *uuid.UUID `json:"id"`
	CreatedAt   *time.Time `json:"createdAt"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Price       *int64     `json:"price"`
	Category    *string    `json:"category"`
	ImagesJson  *any       `json:"imagesJson"`
	IsActive    *bool      `json:"isActive"`
}
