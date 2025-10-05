package stockitems

import (
	httpresponse "api-estoque/internal/model/http_response"
	stockitemsModel "api-estoque/internal/model/stock_items"
	"api-estoque/internal/model/stock_items/response/create"
	getbyid "api-estoque/internal/model/stock_items/response/get_by_id"
	"api-estoque/internal/model/stock_items/response/list"
	stockitemsRepo "api-estoque/internal/repositories/stock_items"
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
		s.Logger.Errorf("(StockItems) List - %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar itens do estoque",
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
		s.Logger.Errorf("(StockItems) Create - %v", err)
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar criacao de item de estoque",
		}
	}

	return &create.CreateResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		ProductId:   *stockItems.ProductId,
		WarehouseId: *stockItems.WarehouseId,
	}
}

func (s *Service) GetByID(idWarehouse *uuid.UUID, idProduct *uuid.UUID) *getbyid.GetByIdResponse {
	stockItems, err := s.Repository.GetByID(idWarehouse, idProduct)
	if err != nil {
		s.Logger.Errorf("(StockItems) GetByID - %v", err)
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de item de estoque por id",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		ProductId:   *stockItems.ProductId,
		WarehouseId: *stockItems.WarehouseId,
		Quantity:    *stockItems.Quantity,
		Reserved:    *stockItems.Reserved,
		UpdatedAt:   *stockItems.UpdatedAt,
	}
}

func (s *Service) Update(stockItems *stockitemsModel.StockItems) *httpresponse.Response {
	err := s.Repository.Update(stockItems)
	if err != nil {
		s.Logger.Errorf("(StockItems) Update - %v", err)
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar atualizacao de estoque",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}

func (s *Service) DeductQuantity(baixa *stockitemsModel.StockItemsBaixa) *httpresponse.Response {
	err := s.Repository.DeductQuantity(baixa)
	if err != nil {
		s.Logger.Errorf("(StockItems) DeductQuantity - %v", err)
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar a dedução de quantidade do estoque",
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
		s.Logger.Errorf("(StockItems) Delete - %v", err)
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao deletar item do estoque",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}
