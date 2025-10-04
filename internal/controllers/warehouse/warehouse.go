package warehouse

import (
	httpresponse "api-estoque/internal/model/http_response"
	warehouseModel "api-estoque/internal/model/warehouse"
	warehouseSrvc "api-estoque/internal/services/warehouse"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *warehouseSrvc.Service
	Logger  *logrus.Logger
}

func New(service *warehouseSrvc.Service, logger *logrus.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

// List godoc
// @Summary Listar armazéns
// @Description Retorna a lista de todos os armazéns cadastrados
// @Tags warehouse
// @Produce json
// @Success 200 {object} httpresponse.Response
// @Failure 500 {object} httpresponse.Response
// @Router /warehouses [get]
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Warehouse) List - req recebida")

	res := c.Service.List()

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Create godoc
// @Summary Criar armazém
// @Description Cria um novo armazém no sistema
// @Tags warehouse
// @Accept json
// @Produce json
// @Param warehouse body warehouseModel.Warehouse true "Warehouse"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 500 {object} httpresponse.Response
// @Router /warehouses [post]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Warehouse) Create - req recebida")

	var warehouse warehouseModel.Warehouse

	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	err = warehouse.ValidateCreate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Create(&warehouse)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// GetByID godoc
// @Summary Buscar armazém por ID
// @Description Retorna um armazém específico pelo seu UUID
// @Tags warehouse
// @Produce json
// @Param id path string true "UUID do Armazém"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /warehouses/{id} [get]
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Warehouse) GetByID - req recebida")

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

// Update godoc
// @Summary Atualizar armazém
// @Description Atualiza os dados de um armazém existente
// @Tags warehouse
// @Accept json
// @Produce json
// @Param warehouse body warehouseModel.Warehouse true "Warehouse"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /warehouses [put]
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Warehouse) Update - req recebida")

	var warehouse warehouseModel.Warehouse

	err := json.NewDecoder(r.Body).Decode(&warehouse)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request invalido, falha ao decodificar body")
		return
	}

	err = warehouse.ValidateUpdate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Update(&warehouse)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Delete godoc
// @Summary Remover armazém
// @Description Exclui um armazém específico pelo seu UUID
// @Tags warehouse
// @Produce json
// @Param id path string true "UUID do Armazém"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /warehouses/{id} [delete]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Warehouse) Delete - req recebida")

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
