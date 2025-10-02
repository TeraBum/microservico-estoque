package warehouse

import (
	"api-estoque/internal/service/warehouse"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Controller struct {
	Service *warehouse.Service
	Logger  *logrus.Logger
}

func New(service *warehouse.Service, logger *logrus.Logger) *Controller {
	return &Controller{
		Service: service,
		Logger:  logger,
	}
}

func (c *Controller) List(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) {

}

func (c *Controller) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id precisa ser um inteiro", http.StatusBadRequest)
		return
	}

	c.Logger.Info(id)
}

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id precisa ser um inteiro", http.StatusBadRequest)
		return
	}

	c.Logger.Info(id)
}

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id precisa ser um inteiro", http.StatusBadRequest)
		return
	}

	c.Logger.Info(id)
}
