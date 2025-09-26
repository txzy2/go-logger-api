package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/txzy2/go-logger-api/pkg"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type Database struct {
	GORM *gorm.DB
}

func NewDatabase(cfg *Config) (*Database, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	// Создаем GORM подключение
	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open gorm database: %w", err)
	} else {
		log.Println("GORM connected successfully")
	}

	return &Database{GORM: gormDB}, nil
}

func NewConfigFromEnv() *Config {
	return &Config{
		Host:     pkg.GetEnv("DB_HOST", "localhost"),
		Port:     pkg.GetEnv("DB_PORT", "5432"),
		User:     pkg.GetEnv("DB_USER", "postgres"),
		Password: pkg.GetEnv("DB_PASS", "password"),
		DBName:   pkg.GetEnv("DB_NAME", "logger_db"),
		SSLMode:  pkg.GetEnv("DB_SSLMODE", "disable"),
	}
}

func (d *Database) Close() error {
	var err error
	if d.GORM != nil {
		if sqlDB, err := d.GORM.DB(); err == nil {
			err = sqlDB.Close()
		}
	}
	return err
}
