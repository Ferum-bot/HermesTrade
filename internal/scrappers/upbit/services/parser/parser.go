package parser

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/scrappers/upbit/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/google/uuid"
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
	const maxAssetsCount = 100
	const assetsPerRequestCount = 10
	const requestsCount = maxAssetsCount/assetsPerRequestCount + 1

	result := make([]model.AssetCurrencyPair, 0, maxAssetsCount)

	currentOffset := int64(0)
	for i := 0; i < requestsCount; i++ {
		_, err := parser.exchangeClient.GetAssetPairs(ctx, "filter", currentOffset, assetsPerRequestCount)
		if err != nil {
			return nil, errors.Wrap(err, "parser.exchangeClient.GetAssetPairs")
		}

		result = append(result, model.AssetCurrencyPair{
			Identifier: model.AssetCurrencyPairIdentifier(uuid.New().String()),
		})
	}

	return result, nil
}
