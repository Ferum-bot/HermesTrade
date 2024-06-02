package main

import (
	"context"
	"errors"
	"fmt"
	assets_storage "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/clients/assets-storage"
	"github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/platform/config"
	found_spreads "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/producers/found-spreads"
	assets_retriever "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/services/assets-retriever"
	spread_hunter "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/services/spread-hunter"
	spread_hunter2 "github.com/Ferum-Bot/HermesTrade/internal/asset-spread-hunter/workers/spread-hunter"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	default_sync_spread_hunter "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/service/default-sync-spread-hunter"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const applicationName = "AssetSpreadHunter"
const metricsServerPort = "8183"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()
	conf := config.NewConfig()

	log.Info("Asset-Spread-Hunter is starting")

	kafkaWriter := configureKafkaWriter(log, conf)
	metricsServer := configureMetricsServer()

	foundSpreadsProducer := found_spreads.NewProducer(kafkaWriter)

	assetsStorageClient := assets_storage.NewClient()
	assetsRetriever := assets_retriever.NewService(log, assetsStorageClient)

	spreadHunterAlgorithm := default_sync_spread_hunter.NewDefaultSyncSpreadHunter()
	spreadHunter := spread_hunter.NewService(spreadHunterAlgorithm)

	spreadHunterWorker := spread_hunter2.NewWorker(log, assetsRetriever, spreadHunter, foundSpreadsProducer)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		log.Infof("Asset-Spread-Hunter worker started")

		err := spreadHunterWorker.Start(ctx)
		if err != nil {
			log.Errorf("Asset-Spread-Hunter worker exited with error: %s", err)
			close(done)
		}
	}()

	go func() {
		log.Infof("Asset-Spread-Hunter metrics server started")

		err := metricsServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server.ListerAndServer: %s", err)
			close(done)
		}
	}()

	<-done
	log.Infof("Asset-Spread-Hunter is stopping")

	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	err = metricsServer.Shutdown(ctx)
	if err != nil {
		log.Errorf("server.Shutdown: %s", err)
		os.Exit(1)
	}

	err = kafkaWriter.Close()
	if err != nil {
		log.Errorf("kafkaWriter.Close: %s", err)
		os.Exit(1)
	}

	log.Infof("Asset-Spread-Hunter stopped")
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

func configureMetricsServer() *http.Server {
	router := chi.NewRouter()
	router.Handle("/metrics", promhttp.Handler())

	server := http.Server{
		Addr:    fmt.Sprintf("localhost:%s", metricsServerPort),
		Handler: router,
	}

	return &server
}

func configureKafkaWriter(
	logger logger.Logger,
	conf *config.Config,
) *kafka.Writer {
	topic, err := conf.GetKafkaTopicFoundSpreads()
	if err != nil {
		logger.Errorf("conf.GetKafkaTopicFoundSpreads: %s", err)
		os.Exit(1)
	}

	kafkaUrl, err := conf.GetKafkaUrl()
	if err != nil {
		logger.Errorf("config.GetKafkaUrl: %s", err)
		os.Exit(1)
	}

	writer := kafka.Writer{
		Addr:                   kafka.TCP(kafkaUrl),
		Topic:                  topic,
		MaxAttempts:            2,
		ReadTimeout:            2 * time.Second,
		WriteTimeout:           2 * time.Second,
		RequiredAcks:           kafka.RequireOne,
		Async:                  false,
		Logger:                 kafka.LoggerFunc(logger.Infof),
		ErrorLogger:            kafka.LoggerFunc(logger.Errorf),
		AllowAutoTopicCreation: true,
	}

	return &writer
}
