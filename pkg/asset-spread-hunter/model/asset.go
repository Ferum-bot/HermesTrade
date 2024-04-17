package model

type AssetUniversalIdentifier int64
type AssetExternalIdentifier int64
type AssetSourceIdentifier int64

type Asset struct {
	UniversalIdentifier AssetUniversalIdentifier
	ExternalIdentifier  AssetExternalIdentifier
	SourceIdentifier    AssetSourceIdentifier
}

type AssetCurrencyPair struct {
	BaseAsset   Asset
	QuotedAsset Asset

	CurrencyRatio AssetsCurrencyRatio
}

type AssetsCurrencyRatio struct {
	Precision int64
	Value     int64
}
