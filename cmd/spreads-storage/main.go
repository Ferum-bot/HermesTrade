package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	get_spreads "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/api/get-spreads"
	save_spreads "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/api/save-spreads"
	search_spreads "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/api/search-spreads"
	found_spreads2 "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/consumers/found-spreads"
	dto "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/generated/schema"
	config2 "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/platform/config"
	spread_link_builder "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/services/spread-link-builder"
	spreads2 "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/services/spreads"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/storage/spreads"
	found_spreads "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/workers/found-spreads"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const applicationName = "SpreadsStorage"
const serverPort = "8887"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()
	config := config2.NewConfig()

	log.Info("Spreads-Storage is starting")

	kafkaReader := configureKafkaReader(log, config)
	router, mongoClient, foundSpreadsWorker := configureRouter(ctx, log, config, kafkaReader)

	server := http.Server{
		Addr:              fmt.Sprintf("localhost:%s", serverPort),
		Handler:           router,
		ReadTimeout:       2 * time.Second,
		ReadHeaderTimeout: 500 * time.Millisecond,
		WriteTimeout:      5 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		log.Infof("Spreads-Storage started on port: %s", serverPort)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server.ListerAndServer: %s", err)
			close(done)
		}
	}()

	go func() {
		log.Info("Spreads-Storage worker started")

		err := foundSpreadsWorker.Run(ctx)
		if err != nil {
			log.Errorf("Worker returned error: %s", err)
			close(done)
		}
	}()

	<-done
	log.Infof("Spreads-Storage is stopping")

	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)

	err = server.Shutdown(ctx)
	if err != nil {
		log.Errorf("server.Shutdown: %s", err)
		os.Exit(1)
	}

	err = mongoClient.Disconnect(ctx)
	if err != nil {
		log.Errorf("mongoClient.Disconnect: %s", err)
		os.Exit(1)
	}

	err = kafkaReader.Close()
	if err != nil {
		log.Errorf("kafkaWriter.Close: %s", err)
		os.Exit(1)
	}

	log.Infof("Spreads-Storage stopped")
}

func configureRouter(
	ctx context.Context,
	logger logger.Logger,
	config config2.SpreadsStorage,
	kafkaReader *kafka.Reader,
) (*chi.Mux, *mongo.Client, *found_spreads.Worker) {
	mongoUrl, err := config.GetMongoUrl()
	if err != nil {
		logger.Errorf("config.GetMongoUrl: %s", err)
		os.Exit(1)
	}

	mongodbLogLevel := options2.LogLevelInfo
	mongoLogOptions := options2.Logger().SetComponentLevel(options2.LogComponentAll, mongodbLogLevel)
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
		logger.Errorf("Can't connect to mongo: %s", err)
		os.Exit(1)
	}

	databaseName, err := config.GetMongoDatabase()
	if err != nil {
		logger.Errorf("config.GetMongoDatabase: %s", err)
		os.Exit(1)
	}

	router := chi.NewRouter()

	mongoDatabase := mongoClient.Database(databaseName)

	spreadsStorage := spreads.NewStorage(mongoDatabase)

	spreadLinkBuilder := spread_link_builder.NewService()
	spreadService := spreads2.NewService(spreadsStorage, spreadLinkBuilder)

	getSpreadsHandler := get_spreads.New(logger, spreadService)
	saveSpreadsHandler := save_spreads.New(logger, spreadService)
	searchSpreadsHandler := search_spreads.New(logger, spreadService)

	server := server{
		getSpreadsHandler:    getSpreadsHandler,
		saveSpreadsHandler:   saveSpreadsHandler,
		searchSpreadsHandler: searchSpreadsHandler,
	}

	router.Handle("/metrics", promhttp.Handler())

	dto.HandlerFromMux(&server, router)

	foundSpreadsConsumer := found_spreads2.NewConsumer(spreadService)
	foundSpreadsWorker := found_spreads.NewWorker(logger, foundSpreadsConsumer, kafkaReader)

	return router, mongoClient, foundSpreadsWorker
}

func configureKafkaReader(
	logger logger.Logger,
	config config2.SpreadsStorage,
) *kafka.Reader {
	topic, err := config.GetKafkaTopicFoundSpreads()
	if err != nil {
		logger.Errorf("config.GetKafkaTopicFoundSpreads: %s", err)
		os.Exit(1)
	}

	kafkaUrl, err := config.GetKafkaUrl()
	if err != nil {
		logger.Errorf("config.GetKafkaUrl: %s", err)
		os.Exit(1)
	}

	consumerGroup, err := config.GetConsumerGroup()
	if err != nil {
		logger.Errorf("config.GetConsumerGroup: %s", err)
		os.Exit(1)
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:           []string{kafkaUrl},
		GroupID:           consumerGroup,
		Topic:             topic,
		MaxWait:           5 * time.Second,
		ReadBatchTimeout:  2 * time.Second,
		HeartbeatInterval: 1 * time.Second,
		SessionTimeout:    10 * time.Second,
		StartOffset:       kafka.FirstOffset,
		Logger:            kafka.LoggerFunc(logger.Infof),
		ErrorLogger:       kafka.LoggerFunc(logger.Errorf),
		IsolationLevel:    kafka.ReadCommitted,
		MaxAttempts:       2,
	})

	return reader
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

type server struct {
	getSpreadsHandler    *get_spreads.Handler
	saveSpreadsHandler   *save_spreads.Handler
	searchSpreadsHandler *search_spreads.Handler
}

func (s *server) PostSpreadsStorageApiV1GetSpreads(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.getSpreadsHandler.GetSpreads(w, r)
}

func (s *server) PutSpreadsStorageApiV1SaveSpreads(
	w http.ResponseWriter,
	r *http.Request,
) {
	s.saveSpreadsHandler.SaveSpreads(w, r)
}

func (s *server) PostSpreadsStorageApiV1SearchSpreads(
	w http.ResponseWriter,
	r *http.Request,
	params dto.PostSpreadsStorageApiV1SearchSpreadsParams,
) {
	s.searchSpreadsHandler.SearchSpreads(w, r, params)
}
