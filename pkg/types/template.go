package types

type TemplateData struct {
	Title string
	Items []AdditionalField
}

// TemplateDataRequest представляет структуру запроса для шаблона
type TemplateDataRequest struct {
	AdditionalFields []AdditionalField `json:"additionalFields" example:"[{key: \"username\", value: \"John\"}, {key: \"reason\", value: \"нарушение правил\"}]"`
}
