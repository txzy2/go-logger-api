package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/internal/repository"
	"github.com/txzy2/go-logger-api/pkg/basic"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type IncidentMiddleware struct {
	repo *repository.Repository
	basic.BaseController[any]
	logger *zap.Logger
}

func NewIncidentMiddleware(repo *repository.Repository, logger *zap.Logger) *IncidentMiddleware {
	return &IncidentMiddleware{
		repo:   repo,
		logger: logger,
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
			m.logger.Error("Error while checking service", zap.Error(err))
			m.Error(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		// Сохраняем данные в контексте для использования в контроллере
		c.Set("incidentData", incidentData)
		c.Next()
	}
}
