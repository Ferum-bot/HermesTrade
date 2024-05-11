package model

import "time"

type SpreadIdentifier string
type SpreadLength int64

type Spread struct {
	Identifier      SpreadIdentifier
	Head            SpreadElement
	MetaInformation SpreadMetaInformation
}

type SpreadElement struct {
	AssetPair AssetCurrencyPair

	NextElement *SpreadElement
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
