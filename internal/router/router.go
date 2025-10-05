package router

import (
	_ "api-estoque/docs"
	"api-estoque/internal/controllers"
	"api-estoque/internal/controllers/product"
	stockitems "api-estoque/internal/controllers/stock_items"
	stockmoves "api-estoque/internal/controllers/stock_moves"
	"api-estoque/internal/controllers/warehouse"
	middleware "api-estoque/internal/middleware/auth"
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
	ProductController    *product.Controller
}

func New(logger *logrus.Logger, controllers *controllers.Controllers) *Router {
	return &Router{
		Logger:               logger,
		Router:               mux.NewRouter(),
		WarehouseController:  controllers.WarehouseController,
		StockItemsController: controllers.StockItemsController,
		StockMovesController: controllers.StockMovesController,
		ProductController:    controllers.ProductController,
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
	r.AttachProductRoutes()
	r.Router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}

func (r *Router) AttachStockItemsRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/estoque/stock-items").Subrouter()

	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockItemsController.List))).Methods(http.MethodGet)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockItemsController.Create))).Methods(http.MethodPost)
	subrouter.Handle("/baixa", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockItemsController.DeductQuantity))).Methods(http.MethodPost)
	subrouter.Handle("/{idWarehouse}/{idProduct}", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockItemsController.GetByID))).Methods(http.MethodGet)
	subrouter.Handle("/{idWarehouse}/{idProduct}", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.StockItemsController.Update))).Methods(http.MethodPut)
	subrouter.Handle("/{idWarehouse}/{idProduct}", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.StockItemsController.Delete))).Methods(http.MethodDelete)
}

func (r *Router) AttachStockMovesRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/estoque/stock-move").Subrouter()

	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockMovesController.List))).Methods(http.MethodGet)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockMovesController.Create))).Methods(http.MethodPost)
	subrouter.Handle("/{id}", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.StockMovesController.GetByID))).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-product/{idProduct}", r.StockMovesController.ListByProduct).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-warehouse/{idWarehouse}", r.StockMovesController.ListByWarehouse).Methods(http.MethodGet)
	subrouter.HandleFunc("/by-warehouse-product{idWarehouse}/{idProduct}", r.StockMovesController.ListByWarehouseAndProduct).Methods(http.MethodGet)
}

func (r *Router) AttachWarehouseRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/estoque/warehouses").Subrouter()

	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.WarehouseController.List))).Methods(http.MethodGet)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.WarehouseController.Create))).Methods(http.MethodPost)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.WarehouseController.Update))).Methods(http.MethodPut)
	subrouter.Handle("/{id}", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.WarehouseController.GetByID))).Methods(http.MethodGet)
	subrouter.Handle("/{id}", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.WarehouseController.Delete))).Methods(http.MethodDelete)
}

func (r *Router) AttachProductRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/estoque/products").Subrouter()

	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.ProductController.List))).Methods(http.MethodGet)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.ProductController.Create))).Methods(http.MethodPost)
	subrouter.Handle("", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.ProductController.Update))).Methods(http.MethodPut)
	subrouter.Handle("/{id}", middleware.JWTAuthMiddleware("Administrador", "Manager")(http.HandlerFunc(r.ProductController.GetByID))).Methods(http.MethodGet)
	subrouter.Handle("/{id}", middleware.JWTAuthMiddleware("Administrador")(http.HandlerFunc(r.ProductController.Delete))).Methods(http.MethodDelete)
}
