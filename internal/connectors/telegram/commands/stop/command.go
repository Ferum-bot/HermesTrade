package stop

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

const CommandName = "/stop"

type command struct {
	logger      logger.Logger
	telegramBot telegramBot
	chatService chatService
}

func NewCommand(
	logger logger.Logger,
	telegramBot telegramBot,
	chatService chatService,
) commands.Command {
	return &command{
		logger:      logger,
		telegramBot: telegramBot,
		chatService: chatService,
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
	err := c.chatService.RemoveChat(ctx, chatID)
	if err != nil {
		c.logger.Warnf("[chatID: %d][userID: %d][command: %s][error: %s]", chatID, authorID, CommandName, err)

		err = c.telegramBot.SendMessage(ctx, chatID, failureStoppingBotMessage)
		if err != nil {
			return errors.Wrap(err, "c.telegramBot.SendMessage failureStoppingBotMessage")
		}

		return nil
	}

	err = c.telegramBot.SendMessage(ctx, chatID, successStoppingBotMessage)
	if err != nil {
		return errors.Wrap(err, "c.telegramBot.SendMessage failureStoppingBotMessage")
	}

	return nil
}
