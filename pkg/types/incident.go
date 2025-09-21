package types

type AdditionalField struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type IncidentData struct {
	Level            string            `json:"level" validate:"required"`
	Message          string            `json:"message" validate:"required"`
	Domain           string            `json:"domain" validate:"required"`
	Action           string            `json:"action" validate:"required"`
	Function         string            `json:"function" validate:"required"`
	Service          string            `json:"service" validate:"required"`
	AdditionalFields []AdditionalField `json:"additionalFields" validate:"required"`
	Date             string            `json:"date" validate:"required"`
}
