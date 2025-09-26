package senders

import (
	"strings"

	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type MailSender struct {
	Data         types.IncidentData
	SendTemplate *models.SendTemplate
	logger       *zap.Logger
}

func (ms *MailSender) PrepareIncidentData() *PreparationResult {
	emailStrings := strings.Split(ms.SendTemplate.To, ",")
	cleanedMessage := make([]EmailMessage, 0, len(emailStrings))

	for _, emailStr := range emailStrings {
		email := strings.TrimSpace(emailStr)
		if email == "" {
			continue
		}

		cleanedMessage = append(cleanedMessage, EmailMessage{
			To:      email,
			Subject: "Уведомление",
			Body:    ms.SendTemplate.Template,
			IsHTML:  true,
		})
	}

	return &PreparationResult{
		Emails:  cleanedMessage,
		Channel: MAIL,
	}
}
