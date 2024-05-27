package found_spreads

import (
	"context"
	"encoding/json"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/segmentio/kafka-go"
)

type Worker struct {
	logger      logger.Logger
	consumer    foundSpreadsConsumer
	kafkaReader *kafka.Reader
}

func NewWorker(
	logger logger.Logger,
	consumer foundSpreadsConsumer,
	kafkaReader *kafka.Reader,
) *Worker {
	return &Worker{
		logger:      logger,
		consumer:    consumer,
		kafkaReader: kafkaReader,
	}
}

func (worker *Worker) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			err := worker.work(ctx)
			if err != nil {
				worker.logger.Errorf("Consumer fetch finished with error: %s", err)
			}
		}
	}
}

func (worker *Worker) work(ctx context.Context) error {
	message, err := worker.kafkaReader.ReadMessage(ctx)
	if err != nil {
		return errors.Wrap(err, "worker.kafkaReader.ReadMessage")
	}

	foundSpread, err := worker.parseFoundSpread(message)
	if err != nil {
		commitErr := worker.kafkaReader.CommitMessages(ctx, message)
		if commitErr != nil {
			worker.logger.Errorf("consumer.reader.CommitMessages: %s", commitErr)
		}

		return errors.Wrap(err, "parseLLMResultMessage")
	}

	err = worker.consumer.Consume(ctx, foundSpread)

	if err != nil {
		worker.logger.Errorf("worker.consumer.Consume: %s", err)

		commitErr := worker.kafkaReader.CommitMessages(ctx, message)
		if commitErr != nil {
			return errors.Wrap(commitErr, "worker.kafkaReader.CommitMessages")
		}
	}

	return nil
}

func (worker *Worker) parseFoundSpread(
	message kafka.Message,
) (model.Spread, error) {
	spread := model.Spread{}

	err := json.Unmarshal(message.Value, &spread)
	if err != nil {
		return spread, errors.Wrap(err, "json.Unmarshal")
	}

	return spread, nil
}
