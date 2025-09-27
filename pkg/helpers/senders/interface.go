package senders

import (
	"github.com/txzy2/go-logger-api/internal/models"
	errors "github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type SenderManagerInterface interface {
	PrepareIncidentData() (*PreparationResult, error)
	Send(sendData []EmailMessage) bool
}

type SenderManager struct {
	MailManager MailSender
}

func NewSenderManager(sendMethod string, incidentType *models.IncidentType, incidentData *types.IncidentData, logger *zap.Logger) (SenderManagerInterface, error) {
	channel := Channel(sendMethod)
	if sendMethod == "" || !ValidChannels[channel] {
		return nil, errors.ErrInvalidChannel
	}

	creator, exists := senderFactory[channel]
	if !exists {
		logger.Error(errors.ErrInvalidChannel.Error())
		return nil, errors.ErrInvalidChannel
	}
	return creator(incidentType, incidentData, logger), nil
}

var senderFactory = map[Channel]func(
	*models.IncidentType,
	*types.IncidentData,
	*zap.Logger,
) SenderManagerInterface{
	EMAIL: func(incidentType *models.IncidentType, incidentData *types.IncidentData, logger *zap.Logger) SenderManagerInterface {
		logger.Info("Mail sender created")
		data := DataToSendMail{
			Message:         incidentData.Message,
			AdditionalField: incidentData.AdditionalFields,
		}
		return &MailSender{IncidentType: incidentType, Data: data, logger: logger}
	},
}
