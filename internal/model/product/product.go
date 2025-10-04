package product

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Product struct {
	Id          *uuid.UUID `json:"id"`
	CreatedAt   *time.Time `json:"createdAt"`
	Name        *string    `json:"name"`
	Description *string    `json:"description"`
	Price       *int64     `json:"price"`
	Category    *string    `json:"category"`
	ImagesJson  *any       `json:"imagesJson"`
	IsActive    *bool      `json:"isActive"`
}

func (p *Product) ValidateCreate() error {
	if p.Name == nil || *p.Name == "" {
		return errors.New("atributo 'name' faltando ou vazio")
	}

	if p.Description == nil || *p.Description == "" {
		return errors.New("atributo 'description' faltando ou vazio")
	}

	if p.Price == nil || *p.Price <= 0 {
		return errors.New("atributo 'price' faltando ou inválido")
	}

	if p.Category == nil || *p.Category == "" {
		return errors.New("atributo 'category' faltando ou vazio")
	}

	if p.IsActive == nil {
		return errors.New("atributo 'is_active' faltando")
	}

	if p.CreatedAt != nil {
		return errors.New("atributo 'created_at' é controlado pela API")
	}

	if p.Id != nil {
		return errors.New("atributo 'id' é controlado pela API")
	}

	return nil
}

func (p *Product) ValidateUpdate() error {
	if p.Id == nil {
		return errors.New("atributo 'id' faltando, necessário para atualizar o produto")
	}

	if p.CreatedAt != nil {
		return errors.New("atributo 'created_at' é controlado pela API")
	}

	// must have at least one field to update
	if (p.Name == nil || *p.Name == "") &&
		(p.Description == nil || *p.Description == "") &&
		p.Price == nil &&
		(p.Category == nil || *p.Category == "") &&
		p.ImagesJson == nil &&
		p.IsActive == nil {
		return errors.New("nenhum atributo informado para atualização")
	}

	return nil
}
