package main

import (
	"net/http"
	"net/http/httptest"
	"simple-crm/customer"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Tests happy path of submitting a well-formed GET /customers request
func TestGetCustomersHandler(t *testing.T) {
	customerRepo := customer.NewCustomerMockRepository()
	customerService := customer.NewCustomerService(customerRepo)
	req, err := http.NewRequest("GET", "/customers", nil)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerService.GetAllCustomers)
	handler.ServeHTTP(rr, req)

	// Checks for 200 status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("getCustomers returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests happy path of submitting a well-formed POST /customers request
func TestAddCustomerHandler(t *testing.T) {
	customerRepo := customer.NewCustomerMockRepository()
	customerService := customer.NewCustomerService(customerRepo)
	requestBody := strings.NewReader(`
		{
			"name": "Example Name",
			"role": "Example Role",
			"email": "Example Email",
			"phone": "5550199",
			"contacted": true
		}
	`)

	req, err := http.NewRequest("POST", "/customers", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerService.PostNewCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 201 status code
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("addCustomer returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	// Checks for JSON response
	if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
		t.Errorf("Content-Type does not match: got %v want %v",
			ctype, "application/json")
	}
}

// Tests unhappy path of deleting a user that doesn't exist
func TestDeleteCustomerHandler(t *testing.T) {
	customerRepo := customer.NewCustomerMockRepository()
	customerService := customer.NewCustomerService(customerRepo)
	id := "e7847fee-3a0e-455e-b151-519bdb9851c7"
	req, err := http.NewRequest("DELETE", "/customers/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": id,
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerService.DeleteCustomer)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("deleteCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}

// Tests unhappy path of getting a user that doesn't exist
func TestGetCustomerHandler(t *testing.T) {
	customerRepo := customer.NewCustomerMockRepository()
	customerService := customer.NewCustomerService(customerRepo)
	id := "e7847fee-3a0e-455e-b151-519bdb9851c7"
	req, err := http.NewRequest("GET", "/customers/"+id, nil)

	if err != nil {
		t.Fatal(err)
	}

	vars := map[string]string{
		"id": id,
	}

	req = mux.SetURLVars(req, vars)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(customerService.GetCustomerById)
	handler.ServeHTTP(rr, req)

	// Checks for 404 status code
	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("getCustomer returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
