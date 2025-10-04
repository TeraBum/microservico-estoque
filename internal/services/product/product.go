package product

import (
	httpresponse "api-estoque/internal/model/http_response"
	productModel "api-estoque/internal/model/product"
	"api-estoque/internal/model/product/response/create"
	getbyid "api-estoque/internal/model/product/response/get_by_id"
	"api-estoque/internal/model/product/response/list"
	productRepo "api-estoque/internal/repositories/product"
	"net/http"

	"github.com/gofrs/uuid"
	"github.com/sirupsen/logrus"
)

type Service struct {
	Repository *productRepo.Repository
	Logger     *logrus.Logger
}

func New(repository *productRepo.Repository, logger *logrus.Logger) *Service {
	return &Service{
		Repository: repository,
		Logger:     logger,
	}
}

func (s *Service) List() *list.ListResponse {
	products, err := s.Repository.List()
	if err != nil {
		s.Logger.Errorf("List products: %v", err)
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao listar produtos",
		}
	}

	return &list.ListResponse{
		Status:   http.StatusOK,
		Msg:      "Sucesso",
		Products: products,
	}
}

func (s *Service) Create(p *productModel.Product) *create.CreateResponse {
	id, err := s.Repository.Create(p)
	if err != nil {
		s.Logger.Errorf("Create product: %v", err)
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao criar produto",
		}
	}

	return &create.CreateResponse{
		Status: http.StatusOK,
		Msg:    "Sucesso",
		Id:     *id,
	}
}

func (s *Service) GetByID(id *uuid.UUID) *getbyid.GetByIdResponse {
	product, err := s.Repository.GetByID(id)
	if err != nil {
		s.Logger.Errorf("Get product by ID: %v", err)
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao buscar produto por ID",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		Id:          product.Id,
		CreatedAt:   product.CreatedAt,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		Price:       product.Price,
		ImagesJson:  product.ImagesJson,
		IsActive:    product.IsActive,
	}
}

func (s *Service) Update(p *productModel.Product) *httpresponse.Response {
	err := s.Repository.Update(p)
	if err != nil {
		s.Logger.Errorf("Update product: %v", err)
		return &httpresponse.Response{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao atualizar produto",
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
		s.Logger.Errorf("Delete product: %v", err)
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
