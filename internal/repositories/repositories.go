package repositories

import (
	stockitems "api-estoque/internal/repositories/stock_items"
	stockmoves "api-estoque/internal/repositories/stock_moves"
	"api-estoque/internal/repositories/warehouse"
)

type Repositories struct {
	StockItemsRepository *stockitems.Repository
	StockMovesRepository *stockmoves.Repository
	WarehouseRepository  *warehouse.Repository
}

func InstanciateRepositories() *Repositories {
	return &Repositories{
		StockItemsRepository: stockitems.New(),
		StockMovesRepository: stockmoves.New(),
		WarehouseRepository:  warehouse.New(),
	}
}
