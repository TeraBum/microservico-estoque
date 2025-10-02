package stockitems

import (
	"time"

	"github.com/gofrs/uuid"
)

type StockItems struct {
	ProductId   uuid.UUID `db:"ProductId" json:"product_id"`
	WarehouseId uuid.UUID `db:"WarehouseId" json:"warehouse_id"`
	Quantity    int64     `db:"Quantity" json:"quantity"`
	Reserved    int64     `db:"Reserved" json:"reserved"`
	UpdatedAt   time.Time `db:"UpdatedAt" json:"updated_at"`
}
