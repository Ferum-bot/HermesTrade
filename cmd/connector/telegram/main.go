package main

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/fallback"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_all"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_1"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_20"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_5"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/start"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/stop"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/all_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_1_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_20_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_5_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger := configureLogger()

	fallbackCommand := fallback.NewCommand()
	startCommand := start.NewCommand()
	stopCommand := stop.NewCommand()

	sendAllCommand := send_all.NewCommand()

	sendProfitability1Command := send_profitability_1.NewCommand()
	sendProfitability5Command := send_profitability_5.NewCommand()
	sendProfitability20Command := send_profitability_20.NewCommand()

	allSpreadsWorker := all_spreads.NewWorker()
	profitability1Worker := profitability_1_spreads.NewWorker()
	profitability5Worker := profitability_5_spreads.NewWorker()
	profitability20Worker := profitability_20_spreads.NewWorker()

	logger.Infof("Telegram Connector is starting")

	ctx := context.Background()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		logger.Info("Worker all_spreads started")

		err := allSpreadsWorker.Start(ctx)
		if err != nil {
			logger.Errorf("worker all_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		logger.Info("Worker profitability_1_spreads started")

		err := profitability1Worker.Start(ctx)
		if err != nil {
			logger.Errorf("worker profitability_1_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		logger.Info("Worker profitability_5_spreads started")

		err := profitability5Worker.Start(ctx)
		if err != nil {
			logger.Errorf("worker profitability_5_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		logger.Info("Worker profitability_20_spreads started")

		err := profitability20Worker.Start(ctx)
		if err != nil {
			logger.Errorf("worker profitability_20_spreads returned error: %v", err)
			close(done)
		}
	}()

	<-done
	logger.Infof("Telegram Connector is stopping")
}

func configureLogger() logger.Logger {

}
