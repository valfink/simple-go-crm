package main

import (
	"log/slog"
	"net/http"
	"simple-crm/customer"
	"simple-crm/middleware"

	"github.com/gorilla/mux"
)

func main() {
	slog.Info("Creating router...")
	router := mux.NewRouter()

	slog.Info("Registering middleware...")
	router.Use(middleware.RequestLogger)

	slog.Info("Registering handlers...")
	router.HandleFunc("/customers", customer.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", customer.GetCustomerById).Methods("GET")

	slog.Info("Starting server...")
	http.ListenAndServe(":3000", router)
}
