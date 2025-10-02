package create

import (
	"github.com/gofrs/uuid"
)

type CreateResponse struct {
	Status      int
	Msg         string
	ProductId   uuid.UUID `json:"product_id"`
	WarehouseId uuid.UUID `json:"warehouse_id"`
}
