package router

import (
	"github.com/gorilla/mux"
	"tsisIsa/internal/handlers"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/players", handlers.GetPlayers).Methods("GET")
	r.HandleFunc("/players/{name}", handlers.GetPlayer).Methods("GET")
	r.HandleFunc("/health", handlers.GetHealth).Methods("GET")

	return r
}
