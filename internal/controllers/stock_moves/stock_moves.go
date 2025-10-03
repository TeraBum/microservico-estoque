package stockmoves

import (
	httpresponse "api-estoque/internal/model/http_response"
	stockmoves "api-estoque/internal/model/stock_moves"
	stockmovesSrvc "api-estoque/internal/services/stock_moves"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *stockmovesSrvc.Service
	Logger  *logrus.Logger
}

func New(service *stockmovesSrvc.Service, logger *logrus.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	res := c.Service.List()

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) ListByProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idProductStr := vars["idProduct"]

	idProduct, err := uuid.FromString(idProductStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "idProduct precisa ser um UUID válido")
		return
	}

	res := c.Service.ListByProduct(&idProduct)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) ListByWarehouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idWarehouseStr := vars["idWarehouse"]

	idWarehouse, err := uuid.FromString(idWarehouseStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "idWarehouse precisa ser um UUID válido")
		return
	}

	res := c.Service.ListByWarehouse(&idWarehouse)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) ListByWarehouseAndProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idWarehouseStr := vars["idWarehouse"]
	idProductStr := vars["idProduct"]

	idWarehouse, err := uuid.FromString(idWarehouseStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "idWarehouse precisa ser um UUID válido")
		return
	}
	idProduct, err := uuid.FromString(idProductStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "idProduct precisa ser um UUID válido")
		return
	}

	res := c.Service.ListByWarehouseAndProduct(&idWarehouse, &idProduct)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var stockItems stockmoves.StockMove

	err := json.NewDecoder(r.Body).Decode(&stockItems)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	res := c.Service.Create(&stockItems)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.FromString(idStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "id precisa ser um UUID válido")
		return
	}

	res := c.Service.GetByID(&id)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}
