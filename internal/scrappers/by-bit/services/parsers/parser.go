package parsers

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/by-bit/model"
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
