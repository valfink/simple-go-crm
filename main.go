package main

import (
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	slog.Info("Starting server...")
	http.ListenAndServe(":3000", router)
}
