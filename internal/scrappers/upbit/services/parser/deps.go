package parser

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/upbit/clients/upbit"
)

type exchangeClient interface {
	GetAssetPairs(
		ctx context.Context,
		filter string,
		offset, limit int64,
	) ([]upbit.ExchangeData, error)
}
