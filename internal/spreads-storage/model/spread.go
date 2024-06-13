package model

import (
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"time"
)

type Spread struct {
	Identifier      SpreadIdentifier
	Head            SpreadElement
	MetaInformation SpreadMetaInformation
}

type SpreadWithLink struct {
	Identifier      SpreadIdentifier
	Head            SpreadElementWithLink
	MetaInformation SpreadMetaInformation
}

type SpreadElement struct {
	AssetPair model.AssetCurrencyPair

	NextElement *SpreadElement
}

type SpreadElementWithLink struct {
	AssetPair AssetPairWithLinks

	NextElement *SpreadElementWithLink
}

type SpreadMetaInformation struct {
	Length               SpreadLength
	ProfitabilityPercent SpreadProfitabilityPercent
	CreatedAt            time.Time
}

type SpreadProfitabilityPercent struct {
	Precision int64
	Value     int64
}
