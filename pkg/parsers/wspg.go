package parsers

import (
	"errors"
	"log"
	"strings"

	"github.com/txzy2/go-logger-api/pkg/types"
)

type WSPGParser struct {
	Data types.IncidentData
}

func (p *WSPGParser) Parse(data string) (map[string]string, error) {
	return map[string]string{
		"service": "WSPG",
		"level":   p.Data.Level,
	}, nil
}

func (p *WSPGParser) ParseMessage(message string) (map[string]string, error) {
	if strings.Contains(message, "|") {
		parts := strings.Split(message, "|")
		log.Println("parts: ", parts)
		return map[string]string{"code": parts[0],
			"message": parts[1],
		}, nil

	}

	return map[string]string{}, errors.New("invalid message")
}
