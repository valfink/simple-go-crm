package customer

import (
	"errors"

	"github.com/google/uuid"
)

type CustomerMockRepository struct {
	allCustomers map[uuid.UUID]Customer
}

func NewCustomerMockRepository() *CustomerMockRepository {
	mockCustomers := prepareMockData()
	return &CustomerMockRepository{
		allCustomers: mockCustomers,
	}
}

func (r *CustomerMockRepository) ListAllCustomers() []*Customer {
	customerList := make([]*Customer, len(r.allCustomers))
	i := 0
	for _, c := range r.allCustomers {
		customerList[i] = &c
		i++
	}

	return customerList
}

func (r *CustomerMockRepository) GetCustomerById(id uuid.UUID) (*Customer, error) {
	customer, customerExists := r.allCustomers[id]

	if !customerExists {
		return &customer, errors.New("Could not find customer with the id " + id.String())
	}

	return &customer, nil
}

func (r *CustomerMockRepository) AddCustomer(customer Customer) {
	r.allCustomers[customer.ID] = customer
}

func (r *CustomerMockRepository) UpdateCustomer(customer Customer) {
	r.allCustomers[customer.ID] = customer
}

func (r *CustomerMockRepository) RemoveCustomer(id uuid.UUID) bool {
	_, customerExists := r.allCustomers[id]
	if !customerExists {
		return false
	}

	delete(r.allCustomers, id)
	return true
}
