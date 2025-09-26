package senders

type Channel string

var (
	MAIL Channel = "mail"
	PUSH Channel = "push"
	TG   Channel = "telegram"
)

type EmailMessage struct {
	To      string
	Subject string
	Body    string
	IsHTML  bool
}

type PreparationResult struct {
	Emails  []EmailMessage
	Channel Channel
}
