package stockitems

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type StockItems struct {
	ProductId   *uuid.UUID `db:"ProductId" json:"product_id"`
	WarehouseId *uuid.UUID `db:"WarehouseId" json:"warehouse_id"`
	Quantity    *int64     `db:"Quantity" json:"quantity"`
	Reserved    *int64     `db:"Reserved" json:"reserved"`
	UpdatedAt   *time.Time `db:"UpdatedAt" json:"updated_at"`
}

func (s *StockItems) ValidateCreate() error {
	if s.ProductId == nil {
		return errors.New("atributo 'product_id' faltando")
	}

	if s.WarehouseId == nil {
		return errors.New("atributo 'warehouse_id' faltando")
	}

	if s.Quantity == nil {
		return errors.New("atributo 'quantity' faltando")
	}

	if s.Reserved == nil {
		return errors.New("atributo 'reserved' faltando")
	}

	return nil
}

func (s *StockItems) ValidateUpdate() error {
	if s.ProductId == nil || s.WarehouseId == nil {
		return errors.New("atributo 'product_id' ou 'warehouse_id' faltando, ambos sao necessarios para identificar os dados")
	}

	if s.UpdatedAt != nil {
		return errors.New("atributo 'update_at' é controlado pela api")
	}

	if s.Quantity == nil && s.Reserved == nil {
		return errors.New("atributo 'quantity' e 'reserved' faltando, nada para alterar")
	}

	return nil
}
