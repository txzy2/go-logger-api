package senders

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strings"

	"github.com/txzy2/go-logger-api/internal/models"
	"github.com/txzy2/go-logger-api/pkg"
	"github.com/txzy2/go-logger-api/pkg/types"
	"go.uber.org/zap"
)

type MailSender struct {
	IncidentType *models.IncidentType
	Data         DataToSendMail
	logger       *zap.Logger
}

func (ms *MailSender) prepareMailBody() string {
	ms.logger.Info("GetTemplate data", zap.Any("data", ms.Data.AdditionalField))

	tmpl, err := template.ParseFiles("./storage/templates/mail.html")
	if err != nil {
		ms.logger.Error("Failed to parse template", zap.Error(err))
		return ""
	}

	var body bytes.Buffer
	data := types.TemplateData{
		Title: ms.IncidentType.SendTemplate.Subject,
		Items: ms.Data.AdditionalField,
	}

	if err := tmpl.Execute(&body, data); err != nil {
		ms.logger.Error("Failed to execute template", zap.Error(err))
		return ""
	}
	return body.String()
}

func (ms *MailSender) PrepareIncidentData() (*PreparationResult, error) {
	emailStrings := strings.Split(ms.IncidentType.SendTemplate.To, ",")
	cleanedMessage := make([]EmailMessage, 0, len(emailStrings))

	mail := ms.prepareMailBody()
	if mail == "" {
		return nil, fmt.Errorf("mail is empty")
	}

	for _, emailStr := range emailStrings {
		email := strings.TrimSpace(emailStr)
		if email == "" {
			continue
		}

		cleanedMessage = append(cleanedMessage, EmailMessage{
			To:      email,
			Subject: ms.IncidentType.SendTemplate.Subject,
			Body:    mail,
			IsHTML:  true,
		})
	}

	return &PreparationResult{
		Emails:  cleanedMessage,
		Channel: EMAIL,
	}, nil
}

func (ms *MailSender) generateToken(sendData []EmailMessage) string {
	key := pkg.GetEnv("WS_PG_KEY", "key")

	messagesJSON, err := json.Marshal(sendData)
	if err != nil {
		ms.logger.Error("Failed to marshal messages", zap.Error(err))
		return ""
	}

	ms.logger.Info("Messages JSON for token", zap.String("messages", string(messagesJSON)))

	// Создаем хеш
	h := sha256.New()
	h.Write([]byte(key))
	h.Write(messagesJSON)
	h.Write([]byte(key))

	return hex.EncodeToString(h.Sum(nil))
}

func (ms *MailSender) Send(sendData []EmailMessage) bool {
	ms.logger.Info("Send mail")
	token := ms.generateToken(sendData)
	messagesUrl := pkg.GetEnv("MESSAGES_URL", "")

	data := map[string]any{
		"token":                        token,
		"another_registration_service": "ws-pg",
		"messages":                     sendData,
	}

	ms.logger.Info("Send mail data", zap.Any("data", data))

	if messagesUrl != "" {
		jsonData, err := json.Marshal(data)
		if err != nil {
			ms.logger.Error("Failed to marshal data", zap.Error(err))
			return false
		}

		req, err := http.NewRequest("POST", messagesUrl+"/api/v1/send_mail", bytes.NewBuffer(jsonData))
		if err != nil {
			ms.logger.Error("Failed to create request", zap.Error(err))
			return false
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			ms.logger.Error("Failed to send request", zap.Error(err))
			return false
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		ms.logger.Info("Send mail response", zap.String("response", string(body)))

		var response MessagesApiResponse
		err = json.Unmarshal(body, &response)
		if err != nil {
			ms.logger.Error("Failed to parse response JSON", zap.Error(err))
			return false
		}

		// Проверяем успешность операции
		if !response.Success {
			ms.logger.Warn("API returned error",
				zap.Int("code", response.Error.Code),
				zap.String("message", response.Error.Message))
			return false
		}

		return true
	}

	return false
}
