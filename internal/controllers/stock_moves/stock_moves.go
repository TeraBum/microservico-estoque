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

// List godoc
// @Summary Listar movimentações de estoque
// @Description Retorna todas as movimentações de estoque
// @Tags stock-moves
// @Produce json
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-move [get]
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) List - req recebida")

	res := c.Service.List()

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// ListByProduct godoc
// @Summary Listar movimentações por produto
// @Description Retorna todas as movimentações de estoque de um produto específico
// @Tags stock-moves
// @Produce json
// @Param idProduct path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-move/by-product/{idProduct} [get]
func (c *Controller) ListByProduct(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) ListByProduct - req recebida")

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

// ListByWarehouse godoc
// @Summary Listar movimentações por armazém
// @Description Retorna todas as movimentações de estoque de um armazém específico
// @Tags stock-moves
// @Produce json
// @Param idWarehouse path string true "UUID do Armazém"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-move/by-warehouse/{idWarehouse} [get]
func (c *Controller) ListByWarehouse(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) ListByWarehouse - req recebida")

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

// ListByWarehouseAndProduct godoc
// @Summary Listar movimentações por armazém e produto
// @Description Retorna todas as movimentações de estoque filtradas por armazém e produto
// @Tags stock-moves
// @Produce json
// @Param idWarehouse path string true "UUID do Armazém"
// @Param idProduct path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-move/by-warehouse-product/{idWarehouse}/{idProduct} [get]
func (c *Controller) ListByWarehouseAndProduct(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) ListByWarehouseAndProduct - req recebida")

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

// Create godoc
// @Summary Criar movimentação de estoque
// @Description Cria uma nova movimentação de estoque
// @Tags stock-moves
// @Accept json
// @Produce json
// @Param stockMove body stockmoves.StockMove true "Movimentação de Estoque"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /stock-move [post]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) Create - req recebida")

	var stockMove stockmoves.StockMove

	err := json.NewDecoder(r.Body).Decode(&stockMove)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	err = stockMove.ValidateCreate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Create(&stockMove)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// GetByID godoc
// @Summary Buscar movimentação de estoque por ID
// @Description Retorna uma movimentação de estoque específica pelo seu ID
// @Tags stock-moves
// @Produce json
// @Param id path string true "UUID da Movimentação"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /stock-move/{id} [get]
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(StockMove) GetByID - req recebida")

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
