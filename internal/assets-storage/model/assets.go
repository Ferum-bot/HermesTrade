package model

type AssetCurrencyPair struct {
	Identifier    AssetCurrencyPairIdentifier
	BaseAsset     Asset
	QuotedAsset   Asset
	CurrencyRatio AssetCurrencyRatio
}

type Asset struct {
	SourceIdentifier    AssetSourceIdentifier
	UniversalIdentifier AssetUniversalIdentifier
	ExternalIdentifier  AssetExternalIdentifier
}

type AssetCurrencyRatio struct {
	Value     int64
	Precision int64
}

type AddAssetCurrencyPairData struct {
	BaseAsset     Asset
	QuotedAsset   Asset
	CurrencyRatio AssetCurrencyRatio
}
