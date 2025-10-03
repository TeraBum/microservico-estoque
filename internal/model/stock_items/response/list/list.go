package list

import (
	stockitems "api-estoque/internal/model/stock_items"
)

type ListResponse struct {
	Status     int                      `json:"-"`
	Msg        string                   `json:"-"`
	StockItems *[]stockitems.StockItems `json:"stock_items"`
}
