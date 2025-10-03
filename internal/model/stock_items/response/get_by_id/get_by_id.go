package getbyid

import (
	"time"

	"github.com/gofrs/uuid"
)

type GetByIdResponse struct {
	Status      int       `json:"-"`
	Msg         string    `json:"-"`
	ProductId   uuid.UUID `json:"product_id"`
	WarehouseId uuid.UUID `json:"warehouse_id"`
	Quantity    int64     `json:"quantity"`
	Reserved    int64     `json:"reserved"`
	UpdatedAt   time.Time `json:"updated_at"`
}
