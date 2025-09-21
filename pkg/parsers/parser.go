package parsers

import (
	errors "github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type ParserInterface interface {
	Parse(data string) (map[string]string, error)
	ParseMessage(message string) (map[string]string, error)
}

type Parser struct{}

func NewParser(service Service, data types.IncidentData) (ParserInterface, error) {
	creator, exists := parserFactory[service]
	if !exists {
		return nil, errors.ErrUnknownService
	}
	return creator(data), nil
}

var parserFactory = map[Service]func(types.IncidentData) ParserInterface{
	WSPG: func(data types.IncidentData) ParserInterface {
		return &WSPGParser{Data: data}
	},
	// ADS: func(data types.IncidentData) ParserInterface {
	// 	return &ADSParser{Data: data}
	// },
}
