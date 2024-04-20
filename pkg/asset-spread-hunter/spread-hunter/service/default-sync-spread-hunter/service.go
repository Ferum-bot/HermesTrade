package default_sync_spread_hunter

import (
	"context"
	"errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type defaultSpreadHunter struct {
}

func NewDefaultSyncSpreadHunter() spread_hunter.SyncSpreadHunter {
	return &defaultSpreadHunter{}
}

func (spreadHunter defaultSpreadHunter) SearchSpread(
	ctx context.Context,
	assetPairs []model.AssetCurrencyPair,
	searchSettings model.SpreadSearchSettings,
) ([]model.Spread, error) {
	return nil, errors.New("not implemented yet")
}
