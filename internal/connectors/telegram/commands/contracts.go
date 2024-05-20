package commands

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

type Command interface {
	Name() string

	CommandReceived(
		ctx context.Context,
		chatID model.ChatID,
		authorID model.UserID,
	) error
}

type FallbackCommand interface {
	CommandReceived(
		ctx context.Context,
		chatID model.ChatID,
		authorID model.UserID,
	) error
}
