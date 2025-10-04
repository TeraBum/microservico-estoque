package repositories

import (
	"api-estoque/internal/repositories/product"
	stockitems "api-estoque/internal/repositories/stock_items"
	stockmoves "api-estoque/internal/repositories/stock_moves"
	"api-estoque/internal/repositories/warehouse"
)

type Repositories struct {
	StockItemsRepository *stockitems.Repository
	StockMovesRepository *stockmoves.Repository
	WarehouseRepository  *warehouse.Repository
	ProductRepository    *product.Repository
}

func InstanciateRepositories() *Repositories {
	return &Repositories{
		StockItemsRepository: stockitems.New(),
		StockMovesRepository: stockmoves.New(),
		WarehouseRepository:  warehouse.New(),
		ProductRepository:    product.New(),
	}
}
