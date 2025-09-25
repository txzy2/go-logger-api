package v1

import (
	"github.com/gin-gonic/gin"
)

// Add добаляет новый код инцидента
// @Summary Добавление типа инцидента
// @Description Добавляет новый тип инцидента для фильтрации входящего потока
// @Tags health
// @Accept json
// @Produce json
// @Success 200 {object} types.TestResponse "Успешная проверка"
// @Failure 500 {object} map[string]interface{} "Ошибка подключения к базе данных"
// @Router /health [get]
func (h *Handler) AddType(c *gin.Context) {
	h.BaseController.OK(c, "TYPES IS WORK", nil)
}
