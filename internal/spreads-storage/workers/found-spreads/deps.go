package found_spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type foundSpreadsConsumer interface {
	Consume(
		ctx context.Context,
		foundSpread model.Spread,
	) error
}
