package spread_hunter

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
)

type Worker struct {
	logger               logger.Logger
	assetsRetriever      assetsRetrieverService
	spreadsHunter        spreadsHunterService
	foundSpreadsProducer foundSpreadsProducer
}

func NewWorker(
	logger logger.Logger,
	assetsRetriever assetsRetrieverService,
	spreadsHunter spreadsHunterService,
	foundSpreadsProducer foundSpreadsProducer,
) *Worker {
	return &Worker{
		logger:               logger,
		assetsRetriever:      assetsRetriever,
		spreadsHunter:        spreadsHunter,
		foundSpreadsProducer: foundSpreadsProducer,
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
				worker.logger.Errorf("work finished with error: %s", err)
			}
		}
	}
}

func (worker *Worker) work(ctx context.Context) error {
	newAssetPairs, err := worker.assetsRetriever.RetrieveNewAssets(ctx)
	if err != nil {
		return errors.Wrap(err, "worker.assetsRetriever.RetrieveNewAssets")
	}

	foundSpreads, err := worker.spreadsHunter.FindSpreads(ctx, newAssetPairs)
	if err != nil {
		return errors.Wrap(err, "worker.spreadsHunter.FindSpreads")
	}

	err = worker.foundSpreadsProducer.Produce(ctx, foundSpreads)
	if err != nil {
		return errors.Wrap(err, "worker.foundSpreadsProducer.Produce")
	}

	return nil
}
