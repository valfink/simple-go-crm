package customer

import (
	"github.com/google/uuid"
)

type Customer struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Role      string    `json:"role"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Contacted bool      `json:"contacted"`
}
