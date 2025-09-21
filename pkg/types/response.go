package types

// APIResponse представляет стандартный ответ API
// @Description Универсальная модель ответа API с поддержкой дженериков
type APIResponse[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

// IncidentCreateResponse представляет ответ при создании инцидента
// @Description Ответ API при успешном создании инцидента
type IncidentCreateResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"Works"`
	Data    interface{} `json:"data,omitempty"`
}

// TestResponse
// @Description Ответ проверки подключения к БД
type TestResponse struct {
	Success bool        `json:"success" example:"true"`
	Message string      `json:"message" example:"DATABASE IS CONNECTED"`
	Data    interface{} `json:"data,omitempty"`
}
