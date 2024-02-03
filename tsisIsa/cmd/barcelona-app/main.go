package main

import (
	"log"
	"net/http"
	"tsisIsa/routers"
)

func main() {
	r := router.NewRouter()

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
