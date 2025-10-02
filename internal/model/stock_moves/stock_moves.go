package stockmoves

import (
	"time"

	"github.com/gofrs/uuid"
)

type StockMoves struct {
	Id          uuid.UUID `db:"Id" json:"id"`
	ProductId   uuid.UUID `db:"ProductId" json:"product_id"`
	WarehouseId uuid.UUID `db:"WarehouseId" json:"warehouse_id"`
	QtyMoved    int64     `db:"QtyMoved" json:"qty_moved"`
	Reason      string    `db:"Reason" json:"reason"`
	CreatedAt   time.Time `db:"CreatedAt" json:"created_at"`
}
