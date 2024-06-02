package parsers

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/by-bit/model"
)

type Parser struct {
	exchangeClient exchangeClient
}

func New(
	exchangeClient exchangeClient,
) *Parser {
	return &Parser{
		exchangeClient: exchangeClient,
	}
}

func (parser *Parser) ParseNewAssetsPairs(
	ctx context.Context,
) ([]model.AssetCurrencyPair, error) {
	//TODO implement me
	panic("implement me")
}
