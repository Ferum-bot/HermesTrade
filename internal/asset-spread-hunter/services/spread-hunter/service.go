package spread_hunter

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/pointers"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type Service struct {
	spreadHunterAlgorithm spreadHunterAlgorithm
}

func NewService(
	algorithm spreadHunterAlgorithm,
) *Service {
	return &Service{
		spreadHunterAlgorithm: algorithm,
	}
}

func (service *Service) FindSpreads(
	ctx context.Context,
	assetPairs []model.AssetCurrencyPair,
) ([]model.Spread, error) {
	searchSettings := model.SpreadSearchSettings{
		MaxSpreadLength: pointers.Int64Pointer(10),
		MinSpreadLength: pointers.Int64Pointer(3),
		MinSearchProfitabilityRatio: &model.SpreadProfitabilityPercent{ // 0.1
			Precision: 1,
			Value:     1,
		},
		MaxSearchProfitabilityRatio: &model.SpreadProfitabilityPercent{ // 20
			Precision: 0,
			Value:     20,
		},
	}

	spreads, err := service.spreadHunterAlgorithm.SearchSpread(ctx, assetPairs, searchSettings)
	if err != nil {
		return nil, errors.Wrap(err, "service.spreadHunterAlgorithm.SearchSpread")
	}

	return spreads, nil
}
