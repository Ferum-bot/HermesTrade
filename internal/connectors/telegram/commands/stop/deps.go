package stop

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

type chatService interface {
	RemoveChat(
		ctx context.Context,
		chatID model.ChatID,
	) error
}

type telegramBot interface {
	SendMessage(
		ctx context.Context,
		chatID model.ChatID,
		message string,
	) error
}
