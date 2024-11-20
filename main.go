package main

import (
	"log/slog"
	"net/http"
	"simple-crm/customer"
	"simple-crm/middleware"

	"github.com/gorilla/mux"
)

func main() {
	slog.Info("Creating and injecting dependencies...")
	customerRepo := customer.NewCustomerMockRepository()
	customerService := customer.NewCustomerService(customerRepo)

	slog.Info("Creating router...")
	router := mux.NewRouter()

	slog.Info("Registering middleware...")
	router.Use(middleware.RequestLogger)

	slog.Info("Registering handlers...")
	router.HandleFunc("/", customerService.ServeHomePage).Methods("GET")
	router.HandleFunc("/customers", customerService.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", customerService.GetCustomerById).Methods("GET")
	router.HandleFunc("/customers", customerService.PostNewCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", customerService.PutCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", customerService.DeleteCustomer).Methods("DELETE")

	slog.Info("Starting server...")
	http.ListenAndServe(":3000", router)
}
