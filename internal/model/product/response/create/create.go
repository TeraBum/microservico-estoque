package create

import (
	"time"

	"github.com/gofrs/uuid"
)

type CreateResponse struct {
	Status    int
	Msg       string
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
}
