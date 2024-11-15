package customer

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID
	Name      string
	Role      string
	Email     string
	Phone     string
	Contacted bool
}
