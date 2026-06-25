package notifier

import (
	"context"

	"go-order-solid/internal/domain"
)

// Notifier — общий контракт для любого отправителя сообщений.
// EmailSender, SMSSender, TelegramSender и другие отправители
// могут использоваться взаимозаменяемо.
type Notifier interface {
	Send(ctx context.Context, message domain.Message) error
}
