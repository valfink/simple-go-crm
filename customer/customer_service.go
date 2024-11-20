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
	slog.Error("Error response", "Status", status, "Error", error)
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

func PostNewCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var customerCreateDTO CustomerCreateDTO

	err := json.NewDecoder(r.Body).Decode(&customerCreateDTO)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse request payload")
		return
	}

	newId, err := uuid.NewRandom()
	if err != nil {
		message := "Could not generate uuid"
		slog.Error(message, "Error", err)
		respondWithError(w, http.StatusInternalServerError, message)
		return
	}

	newCustomer := Customer{
		ID:        newId,
		Name:      customerCreateDTO.Name,
		Role:      customerCreateDTO.Role,
		Email:     customerCreateDTO.Email,
		Phone:     customerCreateDTO.Phone,
		Contacted: customerCreateDTO.Contacted,
	}

	AddCustomer(newCustomer)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newCustomer)
}
