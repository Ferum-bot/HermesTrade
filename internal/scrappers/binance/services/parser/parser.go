package parser

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/binance/model"
)

type Parser struct {
}

func New() *Parser {
	return &Parser{}
}

func (parser *Parser) ParseNewAssetsPairs(
	ctx context.Context,
) ([]model.AssetCurrencyPair, error) {
	//TODO implement me
	panic("implement me")
}
