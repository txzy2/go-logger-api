package main

import (
	"log"

	"github.com/txzy2/go-logger-api/config"
	_ "github.com/txzy2/go-logger-api/docs" // Важно: импорт документации
	"github.com/txzy2/go-logger-api/pkg"
)

// TODO: Сделать автоудаление через lifecircle

// @title Logger Go API
// @version 1.0
// @description API для системы логирования инцидентов с поддержкой PostgreSQL
// @host localhost:8080
// @BasePath /api/v1
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
// @schemes http https
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
