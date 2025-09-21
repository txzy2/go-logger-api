package v1

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/pkg/types"
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
	data := c.MustGet("incidentData").(types.IncidentData)

	if err := types.ValidateIncidentData(data); err != nil {
		log.Printf("Validation error: %v", err.Error())
		h.BaseController.Error(c, 400, "Validation error: "+err.Error())
		return
	}

	log.Printf("Processing incident: %s from service %s", data.Message, data.Service)
	h.BaseController.OK(c, "Incident processed successfully", nil)
}
