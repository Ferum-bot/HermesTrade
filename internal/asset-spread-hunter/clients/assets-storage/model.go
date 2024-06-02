package assets_storage

import "time"

type AssetsFilter struct {
	TimeFilter   *AssetsTimeFilter
	SourceFilter *AssetsSourceFilter
	TypeFilter   *AssetsTypeFilter
}

type AssetsTimeFilter struct {
	StartTime *time.Time
	EndTime   *time.Time
}

type AssetsSourceFilter struct {
	SourceIdentifiers []AssetSourceIdentifier
}

type AssetsTypeFilter struct {
	AssetUniversalIdentifiers []AssetUniversalIdentifier
}

type AssetCurrencyPair struct {
	Identifier AssetCurrencyPairIdentifier
}
