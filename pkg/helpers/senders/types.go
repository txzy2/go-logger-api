package senders

import "github.com/txzy2/go-logger-api/pkg/types"

type Channel string

var (
	EMAIL Channel = "email"
	PUSH  Channel = "push"
	TG    Channel = "telegram"
)

var ValidChannels = map[Channel]bool{
	EMAIL: true,
	PUSH:  true,
	TG:    true,
}

type EmailMessage struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
	IsHTML  bool   `json:"isHTML"`
}

type PreparationResult struct {
	Emails  []EmailMessage
	Channel Channel
}

type DataToSendMail struct {
	Message         string
	AdditionalField []types.AdditionalField
}

type MessagesApiResponse struct {
	Success bool `json:"success"`
	Error   struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}
