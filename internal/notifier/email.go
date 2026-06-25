package notifier

import (
	"context"
	"fmt"

	"go-order-solid/internal/domain"
)

type EmailSender struct{}

func NewEmailSender() *EmailSender {
	return &EmailSender{}
}

func (s *EmailSender) Send(_ context.Context, message domain.Message) error {
	fmt.Printf("Email-уведомление отправлено клиенту %s: %s\n", message.To, message.Text)
	return nil
}
