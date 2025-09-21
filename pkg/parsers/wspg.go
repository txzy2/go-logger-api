package parsers

import "github.com/txzy2/go-logger-api/pkg/types"

type WSPGParser struct {
	Data types.IncidentData
}

func (p *WSPGParser) Parse(data string) (map[string]string, error) {
	return map[string]string{
		"service":   "WSPG",
		"level":     p.Data.Level,
		"inputData": data,
	}, nil
}

func (p *WSPGParser) ParseMessage(message string) (map[string]string, error) {
	return map[string]string{
		"service": "WSPG",
		"message": message,
		"domain":  p.Data.Domain,
	}, nil
}
