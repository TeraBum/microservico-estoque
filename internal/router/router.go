package router

import (
	"api-estoque/internal/controller/product"
	stockitems "api-estoque/internal/controller/stock_items"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Router struct {
	Logger               *logrus.Logger
	Router               *mux.Router
	ProductController    *product.Controller
	StockItemsController *stockitems.Controller
}

func New(logger *logrus.Logger, productController *product.Controller, stockItemsController *stockitems.Controller) *Router {
	return &Router{
		Logger:               logger,
		Router:               mux.NewRouter(),
		ProductController:    productController,
		StockItemsController: stockItemsController,
	}
}

func (r *Router) Run() {
	r.Logger.Info("Iniciando rotas.")
	r.AttachRoutes()

	r.Router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		r.Logger.Warnf("Request para path desconhecido: %s", req.URL.Path)
		http.NotFound(w, req)
	})
}

func (r *Router) AttachRoutes() {
	r.AttachStockItemsRoutes()
	r.AttachProductRoutes()
}

func (r *Router) AttachStockItemsRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/stock-items").Subrouter()

	subrouter.HandleFunc("", r.StockItemsController.List).Methods(http.MethodGet)
	subrouter.HandleFunc("", r.StockItemsController.Create).Methods(http.MethodPost)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.GetByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.Update).Methods(http.MethodPut)
	subrouter.HandleFunc("/{idWarehouse}/{idProduct}", r.StockItemsController.Delete).Methods(http.MethodDelete)
}

func (r *Router) AttachProductRoutes() {
	subrouter := r.Router.PathPrefix("/api/v1/product").Subrouter()

	subrouter.HandleFunc("", r.ProductController.List).Methods(http.MethodGet)
	subrouter.HandleFunc("", r.ProductController.Create).Methods(http.MethodPost)
	subrouter.HandleFunc("/{id}", r.ProductController.GetByID).Methods(http.MethodGet)
	subrouter.HandleFunc("/{id}", r.ProductController.Update).Methods(http.MethodPut)
	subrouter.HandleFunc("/{id}", r.ProductController.Delete).Methods(http.MethodDelete)
}
