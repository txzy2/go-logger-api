package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
	v1 "github.com/txzy2/go-logger-api/internal/delivery/http/v1"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/internal/service"
	"github.com/txzy2/go-logger-api/pkg/database"
)

type App struct {
	logger *zap.Logger
}

func NewApp() *App {
	logger := setupLogger()

	logger.Info("Application initialized",
		zap.String("version", "1.0.0"),
		zap.String("environment", getEnv("APP_ENV", "development")),
	)

	return &App{logger: logger}
}

func (a *App) Run(port string) error {
	defer func() {
		if err := a.logger.Sync(); err != nil {
			log.Printf("Warning: failed to sync logger: %v", err)
		}
	}()

	// Устанавливаем режим Gin
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()

	router.Use(gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	dbConfig := database.NewConfigFromEnv()
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		a.logger.Error("Database connection failed", zap.Error(err))
		return err
	}
	a.logger.Info("Database connected successfully")

	repos := repository.NewRepository(a.logger, db)
	services := service.NewService(repos, a.logger)

	// Инициализация handlers
	handler := v1.NewHandler(services, repos, a.logger)
	handler.InitRoutes(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		a.logger.Info("Starting server", zap.String("port", port))
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal("Server listen error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	a.logger.Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		a.logger.Error("Server shutdown error", zap.Error(err))
		return err
	}

	a.logger.Info("Server exited properly")
	return nil
}
