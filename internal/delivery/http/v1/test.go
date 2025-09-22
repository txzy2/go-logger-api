package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Health проверяет состояние приложения и подключение к базе данных
// @Summary Проверка здоровья приложения
// @Description Проверяет доступность приложения и подключение к базе данных
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} types.TestResponse "Успешная проверка"
// @Failure 500 {object} map[string]interface{} "Ошибка подключения к базе данных"
// @Router /health [get]
func (h *Handler) Health(c *gin.Context) {
	err := h.services.TestService.Ping(c)
	if err != nil {
		h.BaseController.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	h.logger.Info("Database is connected. API is healthy.")
	h.BaseController.OK(c, "DATABASE IS CONNECTED", nil)
}
