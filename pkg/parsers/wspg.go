package parsers

import (
	"errors"
	"strings"

	"github.com/txzy2/go-logger-api/pkg/types"
)

type WSPGParser struct {
	Data types.IncidentData
}

func (p *WSPGParser) Parse(data string) (ParserResponse, error) {
	return ParserResponse{
		Service: "WSPG",
		Level:   data,
	}, nil
}

func (p *WSPGParser) ParseMessage(message string) (ParserMessageResponse, error) {
	if strings.Contains(message, "|") {
		parts := strings.Split(message, "|")
		return ParserMessageResponse{Code: parts[0], Message: parts[1]}, nil
	}

	return ParserMessageResponse{}, errors.New("invalid message")
}
