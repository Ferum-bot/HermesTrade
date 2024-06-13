package model

import "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"

type AssetPairWithLinks struct {
	AssetPair  model.AssetCurrencyPair
	SourceLink string
	PairLink   string
}
