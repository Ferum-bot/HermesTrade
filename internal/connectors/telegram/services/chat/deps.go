package chat

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

type chatStorage interface {
	GetChatByChatID(
		ctx context.Context,
		chatID model.ChatID,
	) (*model.Chat, error)

	GetChatsWithProfitability(
		ctx context.Context,
		settingsType model.ProfitabilitySettingsType,
		skipCount, limit int64,
	) ([]model.Chat, error)

	UpdateChatProfitability(
		ctx context.Context,
		chatID model.ChatID,
		newProfitability model.ProfitabilitySettingsType,
	) error

	CreateChat(
		ctx context.Context,
		chat model.Chat,
	) error

	DeleteChat(
		ctx context.Context,
		chatID model.ChatID,
	) error

	CountChatsWithProfitability(
		ctx context.Context,
		profitability model.ProfitabilitySettingsType,
	) (int64, error)
}
