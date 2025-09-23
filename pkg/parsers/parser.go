package parsers

import (
	"time"

	errors "github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/types"
)

type ParserInterface interface {
	Parse(data string) (ParserResponse, error)
	ParseMessage(message string) (ParserMessageResponse, error)
}

type Parser struct{}

func NewParser(service types.Service, data types.IncidentData) (ParserInterface, error) {
	creator, exists := parserFactory[service]
	if !exists {
		return nil, errors.ErrUnknownService
	}
	return creator(data), nil
}

var parserFactory = map[types.Service]func(types.IncidentData) ParserInterface{
	types.WSPG: func(data types.IncidentData) ParserInterface {
		return &WSPGParser{Data: data}
	},
	// ADS: func(data types.IncidentData) ParserInterface {
	// 	return &ADSParser{Data: data}
	// },
}

func FormatDate(date time.Time) string {
	return date.Format("2006-01-02")
}
