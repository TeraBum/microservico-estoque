package services

import (
	"api-estoque/internal/repositories"
	"api-estoque/internal/services/product"
	stockitems "api-estoque/internal/services/stock_items"
	stockmoves "api-estoque/internal/services/stock_moves"
	"api-estoque/internal/services/warehouse"

	"github.com/sirupsen/logrus"
)

type Services struct {
	StockItemsService *stockitems.Service
	StockMovesService *stockmoves.Service
	WarehouseService  *warehouse.Service
	ProductService    *product.Service
}

func InstanciateServices(repositories *repositories.Repositories, logger *logrus.Logger) *Services {
	return &Services{
		StockItemsService: stockitems.New(repositories.StockItemsRepository, logger),
		StockMovesService: stockmoves.New(repositories.StockMovesRepository, logger),
		WarehouseService:  warehouse.New(repositories.WarehouseRepository, logger),
		ProductService:    product.New(repositories.ProductRepository, logger),
	}
}
