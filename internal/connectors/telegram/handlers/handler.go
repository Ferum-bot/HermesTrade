package handlers

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"strings"
)

type DefaultHandler struct {
	logger            logger.Logger
	availableCommands []commands.Command
	fallbackCommand   commands.FallbackCommand
}

func NewDefaultHandler(
	logger logger.Logger,
	availableCommands []commands.Command,
	fallbackCommand commands.FallbackCommand,
) *DefaultHandler {
	return &DefaultHandler{
		logger:            logger,
		availableCommands: availableCommands,
		fallbackCommand:   fallbackCommand,
	}
}

func (handler *DefaultHandler) Handle(
	ctx context.Context,
	bot *bot.Bot,
	update *models.Update,
) {
	if update.Message == nil {
		return
	}
	incomeCommand := update.Message.Text

	for _, command := range handler.availableCommands {
		if strings.Contains(incomeCommand, command.Name()) {
			chatID := model.ChatID(update.Message.Chat.ID)
			authorID := model.UserID(update.Message.From.ID)

			err := command.CommandReceived(ctx, chatID, authorID)
			if err != nil {
				handler.logger.Errorf(
					"[ChatID: %d][AuthodID: %d]Received error during executing command: %s, err: %s",
					chatID, authorID, command.Name(), err,
				)
			}
		}
	}
}
