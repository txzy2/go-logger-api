package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.GET("/health", h.Health)
		api.GET("/ping", h.Ping)
	}
}
