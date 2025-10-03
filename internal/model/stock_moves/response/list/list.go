package list

import (
	stockmoves "api-estoque/internal/model/stock_moves"
)

type ListResponse struct {
	Status     int                     `json:"-"`
	Msg        string                  `json:"-"`
	StockMoves *[]stockmoves.StockMove `json:"stock_moves"`
}
