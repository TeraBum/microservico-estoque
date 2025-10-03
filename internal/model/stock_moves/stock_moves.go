package stockmoves

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type StockMove struct {
	Id          *uuid.UUID `db:"Id" json:"id"`
	ProductId   *uuid.UUID `db:"ProductId" json:"product_id"`
	WarehouseId *uuid.UUID `db:"WarehouseId" json:"warehouse_id"`
	QtyMoved    *int64     `db:"QtyMoved" json:"qty_moved"`
	Reason      *string    `db:"Reason" json:"reason"`
	CreatedAt   *time.Time `db:"CreatedAt" json:"created_at"`
}

func (s *StockMove) ValidateCreate() error {
	if s.ProductId == nil {
		return errors.New("atributo 'product_id' faltando")
	}

	if s.WarehouseId == nil {
		return errors.New("atributo 'warehouse_id' faltando")
	}

	if s.QtyMoved == nil {
		return errors.New("atributo 'qty_moved' faltando")
	}

	if s.Reason == nil {
		return errors.New("atributo 'reason' faltando")
	}

	return nil
}
