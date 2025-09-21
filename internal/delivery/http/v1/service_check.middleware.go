package v1

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/basic"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type IncidentMiddleware struct {
	repo *repository.Repository
	basic.BaseController[any]
}

func NewIncidentMiddleware(repo *repository.Repository) *IncidentMiddleware {
	return &IncidentMiddleware{
		repo: repo,
	}
}

func (m *IncidentMiddleware) ServiceCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var incidentData types.IncidentData

		if err := c.ShouldBindJSON(&incidentData); err != nil {
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
		}

		if incidentData.Service == "" {
			m.Error(c, http.StatusBadRequest, "Service is empty")
			c.Abort()
		}

		serviceName := incidentData.Service
		_, err := m.repo.IncidentRepository.FindByName(serviceName)

		if err != nil {
			log.Printf("Error while checking service: %s", err.Error())
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
		}

		c.Next()
	}
}
