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
}
