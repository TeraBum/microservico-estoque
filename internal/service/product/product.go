package product

import (
	httpresponse "api-estoque/internal/model/http_response"
	productModel "api-estoque/internal/model/product"
	"api-estoque/internal/model/product/response/create"
	getbyid "api-estoque/internal/model/product/response/get_by_id"
	"api-estoque/internal/model/product/response/list"
	productRepo "api-estoque/internal/repository/product"
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
		return &list.ListResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar consulta para listar produtos",
		}
	}

	return &list.ListResponse{
		Status:   http.StatusOK,
		Msg:      "Sucesso",
		Products: products,
	}
}

func (s *Service) Create(product *productModel.Product) *create.CreateResponse {
	product, err := s.Repository.Create(product)
	if err != nil {
		return &create.CreateResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar criacao de produto",
		}
	}

	return &create.CreateResponse{
		Status:    http.StatusOK,
		Msg:       "Sucesso",
		Id:        product.Id,
		CreatedAt: *product.CreatedAt,
	}
}

func (s *Service) GetByID(id *uuid.UUID) *getbyid.GetByIdResponse {
	product, err := s.Repository.GetByID(id)
	if err != nil {
		return &getbyid.GetByIdResponse{
			Status: http.StatusInternalServerError,
			Msg:    "falha ao executar busca de produto por id",
		}
	}

	return &getbyid.GetByIdResponse{
		Status:      http.StatusOK,
		Msg:         "Sucesso",
		Id:          product.Id,
		Name:        product.Name,
		Description: product.Description,
		Category:    product.Category,
		ImagesJson:  product.ImagesJson,
		IsActive:    product.IsActive,
		Price:       product.Price,
		CreatedAt:   *product.CreatedAt,
	}
}

func (s *Service) Update(product *productModel.Product) *httpresponse.Response {
	err := s.Repository.Update(product)
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
			Msg:    "falha ao deletar produto",
		}
	}

	return &httpresponse.Response{
		Status: http.StatusOK,
		Msg:    "Sucesso",
	}
}
