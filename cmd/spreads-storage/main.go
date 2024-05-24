package main

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	config2 "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/platform/config"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
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

	log.Info("Spreads Storage is starting")
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
