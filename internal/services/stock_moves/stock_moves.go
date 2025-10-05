package stockmoves

import (
	stockmovesModel "api-estoque/internal/model/stock_moves"
	"api-estoque/internal/model/stock_moves/response/create"
	getbyid "api-estoque/internal/model/stock_moves/response/get_by_id"
	"api-estoque/internal/model/stock_moves/response/list"
	stockmovesRepo "api-estoque/internal/repositories/stock_moves"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository *stockmovesRepo.Repository
	Logger     *logrus.Logger
}

func New(repository *stockmovesRepo.Repository, logger *logrus.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) List() *list.ListResponse {
	stockMoves, err := s.Repository.List()
	if err != nil {
		s.Logger.Errorf("(StockMoves) List - %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar movimentos de estoque",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		StockMoves: stockMoves,
	}
}

func (s *Service) ListByProduct(idProduct *uuid.UUID) *list.ListResponse {
	stockMoves, err := s.Repository.ListByProduct(idProduct)
	if err != nil {
		s.Logger.Errorf("(StockMoves) ListByProduct - %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar movimentos de estoque por produto",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		StockMoves: stockMoves,
	}
}

func (s *Service) ListByWarehouse(idWarehouse *uuid.UUID) *list.ListResponse {
	stockMoves, err := s.Repository.ListByWarehouse(idWarehouse)
	if err != nil {
		s.Logger.Errorf("(StockMoves) ListByWarehouse - %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar movimentos de estoque por galpao",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		StockMoves: stockMoves,
	}
}

func (s *Service) ListByWarehouseAndProduct(idWarehouse *uuid.UUID, idProduct *uuid.UUID) *list.ListResponse {
	stockMoves, err := s.Repository.ListByWarehouseAndProduct(idWarehouse, idProduct)
	if err != nil {
		s.Logger.Errorf("(StockMoves) ListByWarehouseAndProduct - %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar movimentos de estoque por galpao e produto",
		}
	}

	return &list.ListResponse{
		Status:     http.StatusOK,
		Msg:        "Sucesso",
		StockMoves: stockMoves,
	}
}

func (s *Service) Create(warehouse *stockmovesModel.StockMove) *create.CreateResponse {
	result, err := s.Repository.Create(warehouse)
	if err != nil {
		s.Logger.Errorf("(StockMoves) Create - %v", err)
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar criacao de movimentacao de estoque",
		}
	}

	return &create.CreateResponse{
		Status: http.StatusOK,
		Msg:    "Sucesso",
		Id:     *result.Id,
	}
}

func (s *Service) GetByID(id *uuid.UUID) *getbyid.GetByIdResponse {
	stockMoves, err := s.Repository.GetByID(id)
	if err != nil {
		s.Logger.Errorf("(StockMoves) GetByID - %v", err)
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de movimentacao de estoque por id",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		Id:          *stockMoves.Id,
		ProductId:   *stockMoves.ProductId,
		WarehouseId: *stockMoves.WarehouseId,
		QtyMoved:    *stockMoves.QtyMoved,
		Reason:      *stockMoves.Reason,
		CreatedAt:   *stockMoves.CreatedAt,
	}
}
