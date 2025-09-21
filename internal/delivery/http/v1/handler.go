package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/internal/service"
)

type Handler struct {
	services           *service.Service
	repos              *repository.Repository
	incidentMiddleware *IncidentMiddleware
}

func NewHandler(services *service.Service, repos *repository.Repository) *Handler {
	return &Handler{
		services:           services,
		repos:              repos,
		incidentMiddleware: NewIncidentMiddleware(repos),
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/health", h.Health)
		api.GET("/ping", h.Ping)

		log := api.Group("/log")
		{
			log.POST("/", h.incidentMiddleware.ServiceCheckMiddleware(), h.Create)
		}
	}
}
