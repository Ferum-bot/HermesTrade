package send_profitability_20

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

type chatService interface {
	UpdateChatSettings(
		ctx context.Context,
		chatID model.ChatID,
		newType model.ProfitabilitySettingsType,
	) error
}

type telegramBot interface {
	SendMessage(
		ctx context.Context,
		chatID model.ChatID,
		message string,
	) error
}
