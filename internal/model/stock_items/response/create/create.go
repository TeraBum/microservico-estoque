package create

import (
	"github.com/gofrs/uuid"
)

type CreateResponse struct {
	Status      int       `json:"-"`
	Msg         string    `json:"-"`
	ProductId   uuid.UUID `json:"product_id"`
	WarehouseId uuid.UUID `json:"warehouse_id"`
}
