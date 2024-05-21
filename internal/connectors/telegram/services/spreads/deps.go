package spreads

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type spreadsStorageClient interface {
	SearchSpreads(
		ctx context.Context,
		parameters model2.SpreadParameters,
	) ([]model.Spread, error)
}
