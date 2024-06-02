package parsers

import (
	"context"
	exchange "github.com/Ferum-Bot/HermesTrade/internal/scrappers/by-bit/client/by-bit"
)

type exchangeClient interface {
	GetAssetPairs(
		ctx context.Context,
		filter string,
		offset, limit int64,
	) ([]exchange.ExchangeData, error)
}
