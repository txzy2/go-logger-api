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

// ServiceCheckMiddleware проверяет существование сервиса перед обработкой инцидента
// @Description Middleware для валидации сервиса в запросе создания инцидента
func (m *IncidentMiddleware) ServiceCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var incidentData types.IncidentData

		if err := c.ShouldBindJSON(&incidentData); err != nil {
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		serviceName := incidentData.Service
		if _, err := m.repo.IncidentRepository.FindByName(serviceName); err != nil {
			log.Printf("Error while checking service: %s", err.Error())
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		// Сохраняем данные в контексте для использования в контроллере
		c.Set("incidentData", incidentData)
		c.Next()
	}
}
