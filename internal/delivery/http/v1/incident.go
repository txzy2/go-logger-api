package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Create(c *gin.Context) {
	log.Println("Incident controller is works")
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Log controller is works",
	})
}
