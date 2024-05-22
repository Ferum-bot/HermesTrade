package main

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/client"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/fallback"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_all"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_1"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_20"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/send_profitability_5"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/start"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/commands/stop"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/config"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/handlers"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/services/chat"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/services/message_converter"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/services/spreads"
	chat2 "github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/storage/chat"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/all_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_1_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_20_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/workers/profitability_5_spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/go-telegram/bot"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	logger := configureLogger()
	telegramConnectorConfig := config.NewConfig()

	mongoUrl, err := telegramConnectorConfig.GetMongoUrl()
	if err != nil {
		logger.Errorf("telegramConnectorConfig.GetMongoUrl: %s", err)
		os.Exit(1)
	}

	mongodbLogLevel := options2.LogLevelInfo
	mongoLogOptions := options2.Logger().
		SetComponentLevel(options2.LogComponentAll, mongodbLogLevel)

	options := options2.Client().
		ApplyURI(mongoUrl).
		SetTimeout(time.Second).
		SetAppName("TelegramConnector").
		SetConnectTimeout(10 * time.Second).
		SetMaxConnecting(10).
		SetMinPoolSize(5).
		SetRetryReads(true).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(10 * time.Second).
		SetLoggerOptions(mongoLogOptions)

	mongoClient, err := mongo.Connect(ctx, options)
	if err != nil {
		logger.Errorf("Can't connect to mongo: %s", err)
		os.Exit(1)
	}

	databaseName, err := telegramConnectorConfig.GetMongoDatabase()
	if err != nil {
		logger.Errorf("telegramConnectorConfig.GetMongoDatabase: %s", err)
		os.Exit(1)
	}

	mongoDatabase := mongoClient.Database(databaseName)

	chatStorage := chat2.NewChatStorage(mongoDatabase)
	chatService := chat.NewChatService(chatStorage)

	spreadsService := spreads.NewSpreadsService()
	spreadMessageConverter := message_converter.NewMessageConverter()

	telegramClient := client.NewTelegramClient()

	fallbackCommand := fallback.NewCommand(telegramClient)
	startCommand := start.NewCommand(telegramClient)
	stopCommand := stop.NewCommand(logger, telegramClient, chatService)
	sendAllCommand := send_all.NewCommand(logger, telegramClient, chatService)

	sendProfitability1Command := send_profitability_1.NewCommand(logger, telegramClient, chatService)
	sendProfitability5Command := send_profitability_5.NewCommand(logger, telegramClient, chatService)
	sendProfitability20Command := send_profitability_20.NewCommand(logger, telegramClient, chatService)

	availableCommands := []commands.Command{
		startCommand, stopCommand, sendAllCommand,
		sendProfitability1Command, sendProfitability5Command, sendProfitability20Command,
	}
	messageHandler := handlers.NewDefaultHandler(logger, availableCommands, fallbackCommand)

	allSpreadsWorker := all_spreads.NewWorker(logger, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability1Worker := profitability_1_spreads.NewWorker(logger, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability5Worker := profitability_5_spreads.NewWorker(logger, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability20Worker := profitability_20_spreads.NewWorker(logger, chatService, spreadsService, spreadMessageConverter, telegramClient)

	botToken, err := telegramConnectorConfig.GetToken()
	if err != nil {
		logrus.Errorf("telegramConnectorConfig.GetToken: %s", err)
		os.Exit(1)
	}

	botOptions := []bot.Option{
		bot.WithDefaultHandler(messageHandler.Handle),
	}
	telegramBot, err := bot.New(botToken, botOptions...)
	if err != nil {
		logger.Errorf("bot.New: %s", err)
		os.Exit(1)
	}

	telegramClient.SetTelegramBot(telegramBot)

	logger.Infof("Telegram Connector is starting")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		telegramBot.Start(ctx)

		close(done)
	}()

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

	logger.Infof("Telegram Connector is started")

	<-done

	logger.Infof("Telegram Connector is stopping")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = mongoClient.Disconnect(ctx)
	if err != nil {
		logger.Errorf("mongoClient.Disconnect: %s", err)
		os.Exit(1)
	}

	_, err = telegramBot.Close(ctx)
	if err != nil {
		logger.Errorf("telegramBot.Close: %s", err)
		os.Exit(1)
	}

	logger.Infof("Telgram Connector stopped")
}

func configureLogger() logger.Logger {
	logger := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	logger.SetReportCaller(true)
	logger.SetFormatter(formatter)

	return logger
}
