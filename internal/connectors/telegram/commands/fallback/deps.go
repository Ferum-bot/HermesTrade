package fallback

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

type telegramBot interface {
	SendMessage(
		ctx context.Context,
		chatID model.ChatID,
		message string,
	) error
}
