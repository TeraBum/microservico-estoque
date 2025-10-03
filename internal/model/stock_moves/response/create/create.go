package create

import (
	"github.com/gofrs/uuid"
)

type CreateResponse struct {
	Status int       `json:"-"`
	Msg    string    `json:"-"`
	Id     uuid.UUID `json:"id"`
}
