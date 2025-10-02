package stockitems

import (
	httpresponse "api-estoque/internal/model/http_response"
	stockitemsModel "api-estoque/internal/model/stock_items"
	"api-estoque/internal/model/stock_items/response/create"
	getbyid "api-estoque/internal/model/stock_items/response/get_by_id"
	"api-estoque/internal/model/stock_items/response/list"
	stockitemsRepo "api-estoque/internal/repository/stock_items"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository *stockitemsRepo.Repository
	Logger     *logrus.Logger
}

func New(repository *stockitemsRepo.Repository, logger *logrus.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) List() *list.ListResponse {
	stockItems, err := s.Repository.List()
	if err != nil {
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar produtos",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		StockItems: stockItems,
	}
}

func (s *Service) Create(stockItems *stockitemsModel.StockItems) *create.CreateResponse {
	stockItems, err := s.Repository.Create(stockItems)
	if err != nil {
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar criacao de produto",
		}
	}

	return &create.CreateResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		ProductId:   stockItems.ProductId,
		WarehouseId: stockItems.WarehouseId,
	}
}

func (s *Service) GetByID(idWarehouse *uuid.UUID, idProduct *uuid.UUID) *getbyid.GetByIdResponse {
	stockItems, err := s.Repository.GetByID(idWarehouse, idProduct)
	if err != nil {
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de produto por id",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		ProductId:   stockItems.ProductId,
		WarehouseId: stockItems.WarehouseId,
	}
}

func (s *Service) Update(stockItems *stockitemsModel.StockItems) *httpresponse.Response {
	err := s.Repository.Update(stockItems)
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

func (s *Service) Delete(idWarehouse *uuid.UUID, idProduct *uuid.UUID) *httpresponse.Response {
	err := s.Repository.Delete(idWarehouse, idProduct)
	if err != nil {
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao deletar produto",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}
