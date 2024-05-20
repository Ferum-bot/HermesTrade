package main

import (
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/fallback"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_all"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_1"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_20"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_5"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/start"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/stop"
)

func main() {

}

func initAvailableCommands() []commands.Command {
	startCommand := start.NewCommand()
	stopCommand := stop.NewCommand()

	sendAllCommand := send_all.NewCommand()

	sendProfitability1Command := send_profitability_1.NewCommand()
	sendProfitability5Command := send_profitability_5.NewCommand()
	sendProfitability20Command := send_profitability_20.NewCommand()

	return []commands.Command{
		startCommand,
		stopCommand,
		sendAllCommand,
		sendProfitability1Command,
		sendProfitability5Command,
		sendProfitability20Command,
	}
}

func initFallbackCommand() commands.FallbackCommand {
	command := fallback.NewCommand()

	return command
}
