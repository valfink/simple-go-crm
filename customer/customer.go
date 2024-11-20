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

type CustomerCreateDTO struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Contacted bool   `json:"contacted"`
}

type CustomerRepository interface {
	ListAllCustomers() []*Customer
	GetCustomerById(id uuid.UUID) (*Customer, error)
	AddCustomer(customer Customer)
	UpdateCustomer(customer Customer)
	RemoveCustomer(id uuid.UUID) bool
}
