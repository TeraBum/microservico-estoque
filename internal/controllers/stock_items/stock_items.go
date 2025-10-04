package stockitems

import (
	httpresponse "api-estoque/internal/model/http_response"
	stockitemsModel "api-estoque/internal/model/stock_items"
	stockitemsSrvc "api-estoque/internal/services/stock_items"
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

// List godoc
// @Summary Listar items do estoque
// @Description Pega todos os registros de item de estoque
// @Tags stock-items
// @Produce json
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-items [get]
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockItem) List - req recebida")

	res := c.Service.List()

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Create godoc
// @Summary Cria item de estoque
// @Description Faz a criação de item de estoque
// @Tags stock-items
// @Accept json
// @Produce json
// @Param stockItem body stockitemsModel.StockItems true "Stock Item"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-items [post]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockItem) Create - req recebida")

	var stockItems stockitemsModel.StockItems

	err := json.NewDecoder(r.Body).Decode(&stockItems)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	err = stockItems.ValidateCreate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Create(&stockItems)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// GetByID godoc
// @Summary Buscar item de estoque por ID
// @Description Retorna um item de estoque específico pelo idWarehouse e idProduct
// @Tags stock-items
// @Produce json
// @Param idWarehouse path string true "UUID do Warehouse"
// @Param idProduct path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /stock-items/{idWarehouse}/{idProduct} [get]
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockItem) GetByID - req recebida")

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

	res := c.Service.GetByID(&idWarehouse, &idProduct)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Update godoc
// @Summary Atualizar item de estoque
// @Description Atualiza os dados de um item de estoque existente
// @Tags stock-items
// @Accept json
// @Produce json
// @Param stockItem body stockitemsModel.StockItems true "Stock Item"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /stock-items/{idWarehouse}/{idProduct} [put]
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockItem) Update - req recebida")

	var stockItems stockitemsModel.StockItems

	err := json.NewDecoder(r.Body).Decode(&stockItems)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	err = stockItems.ValidateUpdate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Update(&stockItems)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Delete godoc
// @Summary Remover item de estoque
// @Description Exclui um item de estoque pelo idWarehouse e idProduct
// @Tags stock-items
// @Produce json
// @Param idWarehouse path string true "UUID do Warehouse"
// @Param idProduct path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /stock-items/{idWarehouse}/{idProduct} [delete]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockItem) Delete - req recebida")

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

	res := c.Service.Delete(&idWarehouse, &idProduct)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}
