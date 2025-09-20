package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/database"
)

func main() {
	var (
		action = flag.String("action", "migrate", "Action to perform: migrate, rollback")
	)
	flag.Parse()

	if err := pkg.LoadEnv(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	cfg := database.NewConfigFromEnv()
	log.Println("Database config:", cfg)

	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	switch *action {
	case "migrate":
		if err := models.AutoMigrate(db.GORM); err != nil {
			log.Fatalf("Migration failed: %v", err)
		}
		fmt.Println("✅ Migration completed successfully")

	case "rollback":
		if err := models.DropTables(db.GORM); err != nil {
			log.Fatalf("Rollback failed: %v", err)
		}
		fmt.Println("✅ Rollback completed successfully")

	default:
		fmt.Printf("Unknown action: %s\n", *action)
		fmt.Println("Available actions: migrate, rollback")
		os.Exit(1)
	}
}
