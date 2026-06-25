package notifier

import (
	"context"
	"fmt"

	"go-order-solid/internal/domain"
)

type SMSSender struct{}

func NewSMSSender() *SMSSender {
	return &SMSSender{}
}

func (s *SMSSender) Send(_ context.Context, message domain.Message) error {
	fmt.Printf("SMS-уведомление отправлено клиенту %s: %s\n", message.To, message.Text)
	return nil
}
