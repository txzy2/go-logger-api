package senders

import (
	"github.com/txzy2/go-logger-api/internal/models"
	errors "github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type SenderManagerInterface interface {
	PrepareIncidentData() *PreparationResult
}

type SenderManager struct {
	MailManager MailSender
}

func NewSenderManager(channel Channel, data types.IncidentData, sendTemplate *models.SendTemplate, logger *zap.Logger) (SenderManagerInterface, error) {
	creator, exists := senderFactory[channel]
	if !exists {
		logger.Error(errors.ErrInvalidChannel.Error())
		return nil, errors.ErrInvalidChannel
	}
	return creator(data, sendTemplate, logger), nil
}

var senderFactory = map[Channel]func(
	types.IncidentData,
	*models.SendTemplate,
	*zap.Logger,
) SenderManagerInterface{
	MAIL: func(data types.IncidentData, sendTemplate *models.SendTemplate, logger *zap.Logger) SenderManagerInterface {
		logger.Info("Mail sender created")
		return &MailSender{Data: data, SendTemplate: sendTemplate, logger: logger}
	},
}
