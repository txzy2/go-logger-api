package main

import (
	"log"

	"github.com/txzy2/go-logger-api/config"
	"github.com/txzy2/go-logger-api/pkg"
)

func main() {
	// Загружаем переменные окружения из .env файла
	if err := pkg.LoadEnv(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	app := config.NewApp()

	if err := app.Run("8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
