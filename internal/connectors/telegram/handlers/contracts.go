package handlers

import (
	"context"
	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

type TelegramMessagesHandler interface {
	Handle(ctx context.Context, bot *bot.Bot, update *models.Update)
}
