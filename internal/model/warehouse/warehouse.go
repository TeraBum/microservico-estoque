package warehouse

import (
	"errors"
	"time"

	"github.com/gofrs/uuid"
)

type Warehouse struct {
	Id        *uuid.UUID `db:"Id" json:"id"`
	Name      *string    `db:"Name" json:"name"`
	Location  *string    `db:"Location" json:"location"`
	CreatedAt *time.Time `db:"CreatedAt" json:"created_at,omitempty"`
}

func (w *Warehouse) ValidateCreate() error {
	if w.Name == nil {
		return errors.New("atributo 'name' faltando")
	}
	if w.Location == nil {
		return errors.New("atributo 'location' faltando")
	}
	return nil
}

func (w *Warehouse) ValidateUpdate() error {
	if w.CreatedAt != nil {
		return errors.New("atributo 'created_at' nao pode ser alterado")
	}
	if w.Id == nil {
		return errors.New("atributo 'id' faltando")
	}
	if w.Location == nil && w.Name == nil {
		return errors.New("atributo 'location' e 'name' faltando, nada para alterar")
	}
	if w.Location != nil {
		if *w.Location == "" {
			return errors.New("atributo 'location' nao pode ser vazio")
		}
	}
	if w.Name != nil {
		if *w.Name == "" {
			return errors.New("atributo 'name' nao pode ser vazio")
		}
	}
	return nil
}
