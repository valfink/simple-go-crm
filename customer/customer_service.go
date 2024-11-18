package customer

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func respondWithError(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	error := errors.New(msg)
	json.NewEncoder(w).Encode(error)
	slog.Error("Responding with error", "Status", status, "Error", error)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	customers := FindAllCustomers()

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, idPresent := mux.Vars(r)["id"]
	if !idPresent {
		respondWithError(w, http.StatusBadRequest, "ID not specified")
		return
	}

	uuid, err := uuid.Parse(id)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Not a valid uuid: "+id)
		return
	}

	customer, err := FindCustomerById(uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find customer with id: "+uuid.String())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customer)
}
