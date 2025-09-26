package pkg

import (
	"errors"
)

var (
	ErrUnknownService = errors.New("unknown service")
	ErrInvalidData    = errors.New("invalid incident data")
	ErrInvalidChannel = errors.New("invalid channel")
)
