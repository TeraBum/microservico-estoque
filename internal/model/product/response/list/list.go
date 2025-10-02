package list

import "api-estoque/internal/model/product"

type ListResponse struct {
	Status   int
	Msg      string
	Products *[]product.Product `json:"products"`
}
