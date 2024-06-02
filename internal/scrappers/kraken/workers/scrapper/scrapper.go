package scrapper

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type Worker struct {
	logger              logger.Logger
	assetsStorageSender assetsStorageSender
	exchangeParser      exchangeParser
}

func NewWorker(
	logger logger.Logger,
	assetsStorageSender assetsStorageSender,
	exchangeParser exchangeParser,
) *Worker {
	return &Worker{
		logger:              logger,
		assetsStorageSender: assetsStorageSender,
		exchangeParser:      exchangeParser,
	}
}

func (worker *Worker) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := worker.work(ctx)
			if err != nil {
				worker.logger.Errorf("Worker tick finished with error: %s", err)
			}
		}
	}
}

func (worker *Worker) work(ctx context.Context) error {
	newAssets, err := worker.exchangeParser.ParseNewAssetsPairs(ctx)
	if err != nil {
		return errors.Wrap(err, "worker.exchangeParser.ParseNewAssetsPairs")
	}

	err = worker.assetsStorageSender.SaveNewAssets(ctx, newAssets)
	if err != nil {
		return errors.Wrap(err, "worker.assetsStorageSender.SaveNewAssets")
	}

	return nil
}
