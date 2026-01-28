package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kirshunya-twink/go-ci-playground/internal/handler"
)

func main() {
	r := gin.Default()

	r.GET("/health", handler.Health)
	r.POST("/sum", handler.Sum)

	log.Println("Server starting on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
