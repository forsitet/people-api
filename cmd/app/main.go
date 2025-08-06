package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"people/internal/router"
	"people/internal/server"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env файл не найден или не загружен:", err)
	}
	container, err := NewContainer()
	if err != nil {
		log.Fatalf("Failed to create container: %v", err)
	}

	appRouter := router.NewRouter(container.GetServer())
	handler := appRouter.SetupRoutes()
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	httpServer := server.NewServer(handler, port)

	if err := httpServer.Start(); err != nil {
		if closeErr := container.Close(); closeErr != nil {
			log.Printf("Error closing container: %v", closeErr)
		}
		log.Fatalf("Failed to start server: %v", err)
	}
}
