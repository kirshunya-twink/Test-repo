package handler

import (
	"log"
	"net/http"
)

func Health(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("ok")); err != nil {
		log.Printf("failed to write response: %v", err)
	}
}
