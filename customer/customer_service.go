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
	slog.Warn("Error response", "Status", status, "Error", error)
}

func respondOK(w http.ResponseWriter, body any) {
	respondOkWithStatus(w, http.StatusOK, body)
}

func respondOkWithStatus(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(body)
	slog.Info("OK Response", "Status", status)
}

func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	customers := FindAllCustomers()

	respondOK(w, customers)
}

func GetCustomerById(w http.ResponseWriter, r *http.Request) {
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

	respondOK(w, customer)
}

func PostNewCustomer(w http.ResponseWriter, r *http.Request) {
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

	AddOrUpdateCustomer(newCustomer)

	respondOkWithStatus(w, http.StatusCreated, newCustomer)
}

func PutCustomer(w http.ResponseWriter, r *http.Request) {
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

	_, err = FindCustomerById(uuid)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Could not find customer with id: "+uuid.String())
		return
	}

	var updatedCustomer Customer
	err = json.NewDecoder(r.Body).Decode(&updatedCustomer)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse request payload")
		return
	}

	if uuid != updatedCustomer.ID {
		respondWithError(w, http.StatusBadRequest, "IDs don't match: "+uuid.String()+" / "+updatedCustomer.ID.String())
		return
	}

	AddOrUpdateCustomer(updatedCustomer)

	respondOK(w, updatedCustomer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
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

	customerDeleted := RemoveCustomer(uuid)
	if !customerDeleted {
		respondWithError(w, 404, "Could not find customer with id: "+uuid.String())
		return
	}

	GetAllCustomers(w, r)
}
