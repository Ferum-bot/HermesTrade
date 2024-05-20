package start

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

const CommandName = "/start"

type command struct {
	telegramBot telegramBot
}

func NewCommand(
	telegramBot telegramBot,
) commands.Command {
	return &command{
		telegramBot: telegramBot,
	}
}

func (c command) Name() string {
	return CommandName
}

func (c command) CommandReceived(
	ctx context.Context,
	chatID model.ChatID,
	authorID model.UserID,
) error {
	err := c.telegramBot.SendMessage(ctx, chatID, initialBotMessage)
	if err != nil {
		return errors.Wrap(err, "c.telegramBot.SendMessage initialBotMessage")
	}

	err = c.telegramBot.SendMessage(ctx, chatID, availableCommandsWithDescription)
	if err != nil {
		return errors.Wrap(err, "c.telegramBot.SendMessage availableCommandsWithDescription")
	}

	return nil
}
