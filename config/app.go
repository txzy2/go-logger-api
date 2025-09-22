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

	"github.com/gin-gonic/gin"
	v1 "github.com/txzy2/go-logger-api/internal/delivery/http/v1"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/internal/service"
	"github.com/txzy2/go-logger-api/pkg/database"
)

type App struct{}

func NewApp() *App {
	return &App{}
}

func (a *App) Run(port string) error {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализация подключения к БД
	dbConfig := database.NewConfigFromEnv()
	db, err := database.NewDatabase(dbConfig)
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)

	// Инициализация handlers
	handler := v1.NewHandler(services, repos)
	handler.InitRoutes(router)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	return nil
}
