package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	assets_storage "github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/clients/assets-storage"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/clients/kraken"
	assets_storage_sender "github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/services/assets-storage-sender"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/services/converter"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/services/parser"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/kraken/workers/scrapper"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const applicationName = "KrakenScrapper"
const metricsServerPort = "8183"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()

	log.Info("Kraken Scrapper is starting")

	exchangeClient := kraken.NewClient()
	assetsConverter := converter.New()
	assetsStorageClient := assets_storage.NewClient()

	assetsStorageSender := assets_storage_sender.New(assetsStorageClient, assetsConverter)
	exchangeParser := parser.New(exchangeClient)

	worker := scrapper.NewWorker(log, assetsStorageSender, exchangeParser)

	metricsServer := configureMetricsServer()

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(ctx)

	go func() {
		log.Infof("Kraken Scrapper worker started")

		err := worker.Start(ctx)
		if err != nil {
			log.Errorf("worker.Start: %s", err)
			close(done)
		}
	}()

	go func() {
		log.Infof("Kraken Scrapper metrics server started")

		err := metricsServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Errorf("server.ListerAndServer: %s", err)
			close(done)
		}
	}()

	<-done
	log.Infof("Kraken Scrapper is stopping")

	cancel()

	ctx, cancel = context.WithTimeout(ctx, 5*time.Second)
	err = metricsServer.Shutdown(ctx)
	if err != nil {
		log.Errorf("server.Shutdown: %s", err)
		os.Exit(1)
	}

	log.Infof("Kraken Scrapper stopped")
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
