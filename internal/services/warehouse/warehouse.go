package warehouse

import (
	httpresponse "api-estoque/internal/model/http_response"
	warehouseModel "api-estoque/internal/model/warehouse"
	"api-estoque/internal/model/warehouse/response/create"
	getbyid "api-estoque/internal/model/warehouse/response/get_by_id"
	"api-estoque/internal/model/warehouse/response/list"
	warehouseRepo "api-estoque/internal/repositories/warehouse"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository *warehouseRepo.Repository
	Logger     *logrus.Logger
}

func New(repository *warehouseRepo.Repository, logger *logrus.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) List() *list.ListResponse {
	warehouses, err := s.Repository.List()
	if err != nil {
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar produtos",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		Warehouses: warehouses,
	}
}

func (s *Service) Create(warehouse *warehouseModel.Warehouse) *create.CreateResponse {
	result, err := s.Repository.Create(warehouse)
	if err != nil {
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar criacao de produto",
		}
	}

	return &create.CreateResponse{
		Status: http.StatusOK,
		Msg:    "Sucesso",
		Id:     *result.Id,
	}
}

func (s *Service) GetByID(id *uuid.UUID) *getbyid.GetByIdResponse {
	warehouse, err := s.Repository.GetByID(id)
	if err != nil {
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de produto por id",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:    http.StatusOK,
		Msg:       "Sucesso",
		Id:        *warehouse.Id,
		Name:      *warehouse.Name,
		Location:  *warehouse.Location,
		CreatedAt: warehouse.CreatedAt,
	}
}

func (s *Service) Update(warehouse *warehouseModel.Warehouse) *httpresponse.Response {
	err := s.Repository.Update(warehouse)
	if err != nil {
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de produto por id",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}

func (s *Service) Delete(id *uuid.UUID) *httpresponse.Response {
	err := s.Repository.Delete(id)
	if err != nil {
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao deletar warehouse",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}
