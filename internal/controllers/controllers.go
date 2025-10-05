package controllers

import (
	"api-estoque/internal/controllers/product"
	stockitems "api-estoque/internal/controllers/stock_items"
	stockmoves "api-estoque/internal/controllers/stock_moves"
	"api-estoque/internal/controllers/warehouse"
	"api-estoque/internal/services"

	"github.com/sirupsen/logrus"
)

type Controllers struct {
	StockItemsController *stockitems.Controller
	StockMovesController *stockmoves.Controller
	WarehouseController  *warehouse.Controller
	ProductController    *product.Controller
}

func InstanciateControllers(services *services.Services, logger *logrus.Logger) *Controllers {
	return &Controllers{
		StockItemsController: stockitems.New(services.StockItemsService, services.StockMovesService, logger),
		StockMovesController: stockmoves.New(services.StockMovesService, logger),
		WarehouseController:  warehouse.New(services.WarehouseService, logger),
		ProductController:    product.New(services.ProductService, logger),
	}
}
