package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"simple-crm/customer"

	"github.com/gorilla/mux"
)

func main() {
	customer.PrepareMockData()
	customers := customer.ListAllCustomers()
	slog.Info("Current Customers")
	for n, c := range customers {
		slog.Info("Customer:", fmt.Sprintf("C%d", n), c)
	}
	router := mux.NewRouter()

	slog.Info("Starting server...")
	http.ListenAndServe(":3000", router)
}
