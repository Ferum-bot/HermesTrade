package assets

type assetCurrencyRow struct {
	Identifier       string        `bson:"identifier"`
	SourceIdentifier int64         `bson:"source_identifier"`
	BaseAsset        asset         `bson:"base_asset"`
	QuotedAsset      asset         `bson:"quoted_asset"`
	CurrencyRatio    currencyRatio `bson:"currency_ratio"`
	CreatedAt        string        `bson:"created_at"`
}

type asset struct {
	UniversalIdentifier int64 `bson:"universal_identifier"`
	ExternalIdentifier  int64 `bson:"external_identifier"`
}

type currencyRatio struct {
	Value     int64 `bson:"value"`
	Precision int64 `bson:"precision"`
}
