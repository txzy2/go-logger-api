package v1

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
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

// GetTemplate обрабатывает запрос на получение HTML шаблона
// @Summary Получить HTML шаблон с данными
// @Description Генерирует HTML страницу на основе переданных данных
// @Tags templates
// @Accept json
// @Produce html
// @Param request body types.TemplateDataRequest true "Данные для шаблона"
// @Success 200 {string} string "HTML страница"
// @Failure 400 {object} object "Неверный формат запроса"
// @Failure 500 {object} object "Внутренняя ошибка сервера"
// @Router /template [post]
func (h *Handler) GetTemplate(c *gin.Context) {
	h.logger.Info("GetTemplate")

	var templateData types.TemplateDataRequest
	if err := c.ShouldBindJSON(&templateData); err != nil {
		h.logger.Error("Invalid request body", zap.Error(err))
		h.BaseController.Error(c, http.StatusBadRequest, "Invalid request body: "+err.Error())
		return
	}

	h.logger.Info("GetTemplate data", zap.Any("data", templateData))

	tmpl, err := template.ParseFiles("./storage/templates/mail.html")
	if err != nil {
		h.BaseController.Error(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Header("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.Execute(
		c.Writer,
		types.TemplateData{
			Title: "Аккаунт заблокирован",
			Items: templateData.AdditionalFields,
		}); err != nil {
		h.BaseController.Error(c, http.StatusInternalServerError, err.Error())
		return
	}
}
