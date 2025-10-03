package getbyid

import (
	"time"

	"github.com/gofrs/uuid"
)

type GetByIdResponse struct {
	Status    int        `json:"-"`
	Msg       string     `json:"-"`
	Id        uuid.UUID  `db:"Id" json:"id"`
	Name      string     `db:"Name" json:"name"`
	Location  string     `db:"Location" json:"location"`
	CreatedAt *time.Time `db:"CreatedAt" json:"created_at,omitempty"`
}
