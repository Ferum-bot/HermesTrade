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
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/handlers"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/platform/config"
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

const applicationName = "TelegramConnector"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()
	telegramConnectorConfig := config.NewConfig()

	log.Infof("Telegram Connector is starting")

	mongoUrl, err := telegramConnectorConfig.GetMongoUrl()
	if err != nil {
		log.Errorf("telegramConnectorConfig.GetMongoUrl: %s", err)
		os.Exit(1)
	}

	mongodbLogLevel := options2.LogLevelInfo
	mongoLogOptions := options2.Logger().
		SetComponentLevel(options2.LogComponentAll, mongodbLogLevel)

	options := options2.Client().
		ApplyURI(mongoUrl).
		SetTimeout(time.Second).
		SetAppName(applicationName).
		SetConnectTimeout(10 * time.Second).
		SetMaxConnecting(10).
		SetMinPoolSize(5).
		SetRetryReads(true).
		SetMaxConnIdleTime(30 * time.Second).
		SetServerSelectionTimeout(10 * time.Second).
		SetLoggerOptions(mongoLogOptions)

	mongoClient, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Errorf("Can't connect to mongo: %s", err)
		os.Exit(1)
	}

	databaseName, err := telegramConnectorConfig.GetMongoDatabase()
	if err != nil {
		log.Errorf("telegramConnectorConfig.GetMongoDatabase: %s", err)
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
	stopCommand := stop.NewCommand(log, telegramClient, chatService)
	sendAllCommand := send_all.NewCommand(log, telegramClient, chatService)

	sendProfitability1Command := send_profitability_1.NewCommand(log, telegramClient, chatService)
	sendProfitability5Command := send_profitability_5.NewCommand(log, telegramClient, chatService)
	sendProfitability20Command := send_profitability_20.NewCommand(log, telegramClient, chatService)

	availableCommands := []commands.Command{
		startCommand, stopCommand, sendAllCommand,
		sendProfitability1Command, sendProfitability5Command, sendProfitability20Command,
	}
	messageHandler := handlers.NewDefaultHandler(log, availableCommands, fallbackCommand)

	allSpreadsWorker := all_spreads.NewWorker(log, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability1Worker := profitability_1_spreads.NewWorker(log, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability5Worker := profitability_5_spreads.NewWorker(log, chatService, spreadsService, spreadMessageConverter, telegramClient)
	profitability20Worker := profitability_20_spreads.NewWorker(log, chatService, spreadsService, spreadMessageConverter, telegramClient)

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
		log.Errorf("bot.New: %s", err)
		os.Exit(1)
	}

	telegramClient.SetTelegramBot(telegramBot)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		telegramBot.Start(ctx)

		close(done)
	}()

	go func() {
		log.Info("Worker all_spreads started")

		err := allSpreadsWorker.Start(ctx)
		if err != nil {
			log.Errorf("worker all_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		log.Info("Worker profitability_1_spreads started")

		err := profitability1Worker.Start(ctx)
		if err != nil {
			log.Errorf("worker profitability_1_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		log.Info("Worker profitability_5_spreads started")

		err := profitability5Worker.Start(ctx)
		if err != nil {
			log.Errorf("worker profitability_5_spreads returned error: %v", err)
			close(done)
		}
	}()

	go func() {
		log.Info("Worker profitability_20_spreads started")

		err := profitability20Worker.Start(ctx)
		if err != nil {
			log.Errorf("worker profitability_20_spreads returned error: %v", err)
			close(done)
		}
	}()

	log.Infof("Telegram Connector is started")

	<-done

	log.Infof("Telegram Connector is stopping")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err = mongoClient.Disconnect(ctx)
	if err != nil {
		log.Errorf("mongoClient.Disconnect: %s", err)
		os.Exit(1)
	}

	_, err = telegramBot.Close(ctx)
	if err != nil {
		log.Errorf("telegramBot.Close: %s", err)
		os.Exit(1)
	}

	log.Infof("Telgram Connector stopped")
}

func configureLogger() logger.Logger {
	log := logrus.New()

	formatter := new(logrus.JSONFormatter)
	formatter.TimestampFormat = "2006-01-02 15:04:05.000"
	formatter.PrettyPrint = false

	log.SetReportCaller(true)
	log.SetFormatter(formatter)

	return log.WithFields(logrus.Fields{
		"application": applicationName,
	})
}
