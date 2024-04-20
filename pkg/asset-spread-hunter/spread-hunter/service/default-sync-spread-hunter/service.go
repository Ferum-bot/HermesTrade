package default_sync_spread_hunter

import (
	"errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type service struct {
}

func NewDefaultSyncSpreadHunter() spread_hunter.SyncSpreadHunter {
	return &service{}
}

func (s service) SearchSpread(
	assetPairs []model.AssetCurrencyPair,
	searchSettings model.SpreadSearchSettings,
) ([]model.Spread, error) {
	return nil, errors.New("not implemented yet")
}
