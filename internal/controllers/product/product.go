package product

import (
	httpresponse "api-estoque/internal/model/http_response"
	productModel "api-estoque/internal/model/product"
	productSrvc "api-estoque/internal/services/product"
	"encoding/json"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *productSrvc.Service
	Logger  *logrus.Logger
}

func New(service *productSrvc.Service, logger *logrus.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

// List godoc
// @Summary Listar produtos
// @Description Retorna todos os produtos cadastrados
// @Tags products
// @Produce json
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /products [get]
func (c *Controller) List(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Product) List - req recebida")

	res := c.Service.List()

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Create godoc
// @Summary Criar produto
// @Description Faz a criação de um novo produto
// @Tags products
// @Accept json
// @Produce json
// @Param product body productModel.Product true "Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Router /products [post]
func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Product) Create - req recebida")

	var product productModel.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request inválido, falha ao decodificar body")
		return
	}

	err = product.ValidateCreate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Create(&product)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// GetByID godoc
// @Summary Buscar produto por ID
// @Description Retorna um produto específico pelo seu ID
// @Tags products
// @Produce json
// @Param id path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /products/{id} [get]
func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Product) GetByID - req recebida")

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
// @Summary Atualizar produto
// @Description Atualiza os dados de um produto existente
// @Tags products
// @Accept json
// @Produce json
// @Param product body productModel.Product true "Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /products/{id} [put]
func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Product) Update - req recebida")

	var product productModel.Product

	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, "request inválido, falha ao decodificar body")
		return
	}

	err = product.ValidateUpdate()
	if err != nil {
		httpresponse.JSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	res := c.Service.Update(&product)

	if res.Status != http.StatusOK {
		httpresponse.JSONError(w, res.Status, res.Msg)
		return
	}

	httpresponse.JSONSuccess(w, res)
}

// Delete godoc
// @Summary Remover produto
// @Description Exclui um produto pelo seu ID
// @Tags products
// @Produce json
// @Param id path string true "UUID do Produto"
// @Success 200 {object} httpresponse.Response
// @Failure 400 {object} httpresponse.Response
// @Failure 404 {object} httpresponse.Response
// @Router /products/{id} [delete]
func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	c.Logger.Info("(Product) Delete - req recebida")

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
