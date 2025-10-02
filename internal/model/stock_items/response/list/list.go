package list

import (
	stockitems "api-estoque/internal/model/stock_items"
)

type ListResponse struct {
	Status     int
	Msg        string
	StockItems *[]stockitems.StockItems `json:"stock_items"`
}
