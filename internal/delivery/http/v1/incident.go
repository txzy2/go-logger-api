package v1

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Create создает новый инцидент в системе логирования
// @Summary Создание инцидента
// @Description Создает новый инцидент в системе логирования с проверкой сервиса
// @Tags incident
// @Accept json
// @Produce json
// @Param incident body types.IncidentData true "Данные инцидента"
// @Success 200 {object} types.IncidentCreateResponse "Инцидент успешно создан"
// @Failure 400 {object} map[string]interface{} "Некорректные данные запроса"
// @Failure 500 {object} map[string]interface{} "Внутренняя ошибка сервера"
// @Router /log [post]
func (h *Handler) Create(c *gin.Context) {
	log.Println("Incident controller is works!")

	h.BaseController.OK(c, "Works", nil)
}
