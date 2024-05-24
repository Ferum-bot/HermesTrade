package main

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const applicationName = "CoinbaseScrapper"

func main() {
	err := godotenv.Load()
	if err != nil {
		logrus.Warnf("godotenv.Load: %s", err)
	}

	ctx := context.Background()
	log := configureLogger()

	log.Info("Coinbase Scrapper is starting")
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
