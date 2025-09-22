package v1

import (
	"log"
	"net/http"
	"runtime/debug"

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
// @Failure 500 {object} map[string]interface{} "Ошибка валидации данных"
// @Router /log [post]
func (h *Handler) Create(c *gin.Context) {
	data := c.MustGet("incidentData").(types.IncidentData)

	if err := types.ValidateIncidentData(data); err != nil {
		log.Printf("Validation error: %v", err.Error())
		h.BaseController.Error(c, http.StatusBadRequest, "Invalid data")
		return
	}

	h.BaseController.OK(c, "SUCCESS", nil)

	go h.processIncidentBackground(data)
}

func (h *Handler) processIncidentBackground(data types.IncidentData) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic in background processing: %v", r)
			debug.PrintStack()
		}
	}()

	res := h.services.IncidentService.WriteOrSaveLogs(data)
	log.Printf("Result of parsing: %s", res)
}
