package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health проверяет состояние приложения
func (h *Handler) Health(c *gin.Context) {
	h.services.TestService.Ping(c)
	c.JSON(http.StatusOK, gin.H{
		"status":  "healthy",
		"message": "Application is running",
	})
}

func (h *Handler) Ping(c *gin.Context) {
	err := h.services.TestService.Ping(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to ping",
			"details": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "PONG. DATABASE IS CONNECTED",
	})
}
