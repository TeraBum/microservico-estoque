package stockitems

import (
	httpresponse "api-estoque/internal/model/http_response"
	stockitemsModel "api-estoque/internal/model/stock_items"
	stockitemsSrvc "api-estoque/internal/service/stock_items"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *stockitemsSrvc.Service
	Logger  *logrus.Logger
}

func New(service *stockitemsSrvc.Service, logger *logrus.Logger) *Controller {
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

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	var stockItems stockitemsModel.StockItems

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

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	var stockItems stockitemsModel.StockItems

	err := json.NewDecoder(r.Body).Decode(&stockItems)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	res := c.Service.Update(&stockItems)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := uuid.FromString(idStr)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "id precisa ser um UUID válido")
		return
	}

	res := c.Service.Delete(&id)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}
