package model

import "time"

type AssetFilters struct {
	TimeFilter   *AssetTimeFilter
	SourceFilter *AssetSourceFilter
	TypeFilter   *AssetTypeFilter
}

type AssetTimeFilter struct {
	StartTime time.Time
	EndTime   time.Time
}

type AssetSourceFilter struct {
	SourceIdentifiers []AssetSourceIdentifier
}

type AssetTypeFilter struct {
	UniversalIdentifiers []AssetUniversalIdentifier
}
