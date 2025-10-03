package list

import (
	"api-estoque/internal/model/warehouse"
)

type ListResponse struct {
	Status     int                    `json:"-"`
	Msg        string                 `json:"-"`
	Warehouses *[]warehouse.Warehouse `json:"warehouses"`
}
