package router

import (
	"api-estoque/internal/controllers"
	stockitems "api-estoque/internal/controllers/stock_items"
	stockmoves "api-estoque/internal/controllers/stock_moves"
	"api-estoque/internal/controllers/warehouse"
	_ "api-estoque/internal/docs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Router struct {
	Logger               *logrus.Logger
	Router               *mux.Router
	WarehouseController  *warehouse.Controller
	StockItemsController *stockitems.Controller
	StockMovesController *stockmoves.Controller
}

func New(logger *logrus.Logger, controllers *controllers.Controllers) *Router {
	return &Router{
		Logger:               logger,
		Router:               mux.NewRouter(),
		WarehouseController:  controllers.WarehouseController,
		StockItemsController: controllers.StockItemsController,
		StockMovesController: controllers.StockMovesController,
	}
}

func (r *Router) Run() {
	r.Logger.Info("Iniciando rotas...")
	r.AttachRoutes()

	r.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.Logger.Warnf("Request para path desconhecido: %s", req.URL.Path)
		http.NotFound(w, req)
	})
}

func (r *Router) AttachRoutes() {
	r.AttachStockItemsRoutes()
	r.AttachWarehouseRoutes()
	r.AttachStockMovesRoutes()
	r.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func (r *Router) AttachStockItemsRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/stock-items").Subrouter()

	subrouter.HandleFunc("", r.StockItemsController.List).Methods(http.MethodGet)
	subrouter.HandleFunc("", r.StockItemsController.Create).Methods(http.MethodPost)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.GetByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.Update).Methods(http.MethodPut)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.Delete).Methods(http.MethodDelete)
}

func (r *Router) AttachStockMovesRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/stock-move").Subrouter()

	subrouter.HandleFunc("", r.StockMovesController.List).Methods(http.MethodGet)
	subrouter.HandleFunc("", r.StockMovesController.Create).Methods(http.MethodPost)
	subrouter.HandleFunc("/{id}", r.StockMovesController.GetByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-product/{idProduct}", r.StockMovesController.ListByProduct).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-warehouse/{idWarehouse}", r.StockMovesController.ListByWarehouse).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-warehouse-product{idWarehouse}/{idProduct}", r.StockMovesController.ListByWarehouseAndProduct).Methods(http.MethodGet)
}

func (r *Router) AttachWarehouseRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/warehouses").Subrouter()

	subrouter.HandleFunc("", r.WarehouseController.List).Methods(http.MethodGet)
	subrouter.HandleFunc("", r.WarehouseController.Create).Methods(http.MethodPost)
	subrouter.HandleFunc("", r.WarehouseController.Update).Methods(http.MethodPut)
	subrouter.HandleFunc("/{id}", r.WarehouseController.GetByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", r.WarehouseController.Delete).Methods(http.MethodDelete)
}
