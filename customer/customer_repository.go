package customer

import (
	"errors"

	"github.com/google/uuid"
)

var allCustomers = PrepareMockData()

func FindAllCustomers() []*Customer {
	customerList := make([]*Customer, len(allCustomers))
	i := 0
	for _, c := range allCustomers {
		customerList[i] = &c
		i++
	}

	return customerList
}

func FindCustomerById(id uuid.UUID) (*Customer, error) {
	customer, customerExists := allCustomers[id]

	if !customerExists {
		return &customer, errors.New("Could not find customer with the id " + id.String())
	}

	return &customer, nil
}

func AddOrUpdateCustomer(customer Customer) {
	allCustomers[customer.ID] = customer
}

func RemoveCustomer(id uuid.UUID) bool {
	_, customerExists := allCustomers[id]
	if !customerExists {
		return false
	}

	delete(allCustomers, id)
	return true
}
