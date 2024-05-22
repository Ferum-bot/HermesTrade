package fallback

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type command struct {
	telegramBot telegramBot
}

func NewCommand(
	telegramBot telegramBot,
) commands.FallbackCommand {
	return &command{
		telegramBot: telegramBot,
	}
}

func (command *command) CommandReceived(
	ctx context.Context,
	chatID model.ChatID,
	authorID model.UserID,
) error {
	err := command.telegramBot.SendMessage(ctx, chatID, unKnownCommandMessage)
	if err != nil {
		return errors.Wrap(err, "command.telegramBot.SendMessage")
	}

	return nil
}
