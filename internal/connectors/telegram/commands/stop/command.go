package stop

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
)

const CommandName = "/stop"

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
	//TODO implement me
	panic("implement me")
}
