package main

import (
	"log"
	"net/http"
	"tic-tac-toe/routes"
)

func main() {
	router := routes.SetupRoutes()

	log.Println("Starting server at port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
