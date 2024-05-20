package client

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/go-telegram/bot"
)

type TelegramClient struct {
	bot *bot.Bot
}

func NewTelegramClient(
	bot *bot.Bot,
) *TelegramClient {
	return &TelegramClient{
		bot: bot,
	}
}

func (client *TelegramClient) SendMessage(
	ctx context.Context,
	chatID model.ChatID,
	message string,
) error {
	parameters := bot.SendMessageParams{
		ChatID: chatID,
		Text:   message,
	}
	_, err := client.bot.SendMessage(ctx, &parameters)
	if err != nil {
		return errors.Wrap(err, "client.bot.SendMessage")
	}

	return nil
}
