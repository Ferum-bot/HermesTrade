package main

import (
	"context"
	"errors"
	"fmt"
	add_assets "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/api/add-assets"
	get_assets "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/api/get-assets"
	dto "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/platform/config"
	assets2 "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/services/assets"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/storage/assets"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	options2 "go.mongodb.org/mongo-driver/mongo/options"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const serverPort = "8886"
const applicationName = "AssetsStorage"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()
	conf := config.NewConfig()

	log.Info("Assets-Storage is starting")

	router, mongoClient := configureRouter(ctx, log, conf)

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

	go func() {
		log.Infof("Assets-Storage started on port: %s", serverPort)

		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server.ListerAndServer: %s", err)
			close(done)
		}
	}()

	<-done
	log.Infof("Assets-Storage is stopping")

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

	log.Infof("Assets-Storage stopped")
}

func configureRouter(
	ctx context.Context,
	logger logger.Logger,
	config config.AssetsStorage,
) (*chi.Mux, *mongo.Client) {
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

	mongoDatabase := mongoClient.Database(databaseName)

	assetsStorage := assets.NewAssetsStorage(mongoDatabase)
	assetsService := assets2.NewAssetsService(assetsStorage)

	addAssetsHandler := add_assets.New(logger, assetsService)
	getAssetsHandler := get_assets.New(logger, assetsService)

	router := chi.NewRouter()

	server := server{
		addAssetsHandler: addAssetsHandler,
		getAssetsHandler: getAssetsHandler,
	}

	router.Handle("/metrics", promhttp.Handler())

	dto.HandlerFromMux(&server, router)

	return router, mongoClient
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
	addAssetsHandler *add_assets.Handler
	getAssetsHandler *get_assets.Handler
}

func (s *server) PutAssetsStorageApiV1AddAssets(
	response http.ResponseWriter,
	request *http.Request,
	params dto.PutAssetsStorageApiV1AddAssetsParams,
) {
	s.addAssetsHandler.AddAssets(response, request, params)
}

func (s *server) PostAssetsStorageApiV1GetAssets(
	response http.ResponseWriter,
	request *http.Request,
	params dto.PostAssetsStorageApiV1GetAssetsParams,
) {
	s.getAssetsHandler.GetAssets(response, request, params)
}
