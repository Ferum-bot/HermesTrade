package parser

import (
	"context"
	exchange "github.com/Ferum-Bot/HermesTrade/internal/scrappers/binance/clients/binance"
)

type exchangeClient interface {
	GetAssetPairs(
		ctx context.Context,
		filter string,
		offset, limit int64,
	) ([]exchange.ExchangeData, error)
}
