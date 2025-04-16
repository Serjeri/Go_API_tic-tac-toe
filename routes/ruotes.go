package routes

import (
	"github.com/gorilla/mux"
	"tic-tac-toe/handlers"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/start", handlers.StartNewGame).Methods("POST")
	r.HandleFunc("/move", handlers.Move).Methods("POST")

	return r
}
