package parser

import (
	"context"
	exchange "github.com/Ferum-Bot/HermesTrade/internal/scrappers/coinbase/client/coinbase"
)

type exchangeClient interface {
	GetAssetPairs(
		ctx context.Context,
		filter string,
		offset, limit int64,
	) ([]exchange.ExchangeData, error)
}
