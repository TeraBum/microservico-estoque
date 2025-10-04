package list

import (
	"api-estoque/internal/model/product"
)

type ListResponse struct {
	Status   int                `json:"-"`
	Msg      string             `json:"-"`
	Products *[]product.Product `json:"products"`
}
