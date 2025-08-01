package main

import (
	"log"
	"net/http"
	"os"

	"mailqueue/internal/rabbitmq"
	"mailqueue/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	rabbitmq.Setup(os.Getenv("RABBITMQ_URL"))
	go rabbitmq.StartConsumer()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := routes.SetupRouter()

	log.Printf("Starting server on port: %s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
