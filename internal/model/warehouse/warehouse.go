package warehouse

import (
	"time"

	"github.com/gofrs/uuid"
)

type Warehouse struct {
	Id        uuid.UUID  `db:"Id" json:"id"`
	Name      string     `db:"Name" json:"name"`
	Location  string     `db:"Location" json:"location"`
	CreatedAt *time.Time `db:"CreatedAt" json:"created_at,omitempty"`
}
