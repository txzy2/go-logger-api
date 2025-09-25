package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/internal/service"
	"github.com/txzy2/go-logger-api/pkg/basic"
	"go.uber.org/zap"
)

type Handler struct {
	services           *service.Service
	repos              *repository.Repository
	incidentMiddleware *IncidentMiddleware
	basic.BaseController[any]
	logger *zap.Logger
}

func NewHandler(services *service.Service, repos *repository.Repository, zapLogger *zap.Logger) *Handler {
	return &Handler{
		services:           services,
		repos:              repos,
		incidentMiddleware: NewIncidentMiddleware(repos, zapLogger),
		logger:             zapLogger,
	}
}

// @title Logger Go API
// @version 1.0
// @description API для системы логирования инцидентов
// @host localhost:8080
// @BasePath /api/v1
// @contact.name API Support
// @contact.email support@example.com
// @license.name MIT
// @license.url https://opensource.org/licenses/MIT
func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		test := api.Group("/test")
		{
			test.GET("/health", h.Health)
			test.POST("/template", h.GetTemplate)
		}

		log := api.Group("/log")
		{
			log.POST("/", h.incidentMiddleware.ServiceCheckMiddleware(), h.Create)
		}

		types := api.Group("/types")
		{
			types.POST("/add", h.AddType)
		}
	}
}
