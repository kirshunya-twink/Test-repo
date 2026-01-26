package main

import (
	"log"
	"net/http"

	"github.com/kirshunya-twink/go-ci-playground/internal/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/sum", handler.Sum)

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
