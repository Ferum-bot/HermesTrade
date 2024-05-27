package found_spreads

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type Consumer struct {
	spreadsService spreadsService
}

func NewConsumer(
	spreadsService spreadsService,
) *Consumer {
	return &Consumer{
		spreadsService: spreadsService,
	}
}

func (consumer *Consumer) Consume(
	ctx context.Context,
	foundSpread model.Spread,
) error {
	spread := model2.Spread{
		Identifier: model2.SpreadIdentifier(foundSpread.Identifier),
	}

	_, err := consumer.spreadsService.SaveSpreads(ctx, []model2.Spread{spread})
	if err != nil {
		return errors.Wrap(err, "consumer.spreadsService.SaveSpreads")
	}

	return nil
}
