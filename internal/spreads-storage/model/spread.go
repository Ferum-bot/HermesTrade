package model

type Spread struct {
	Identifier SpreadIdentifier
}

type SpreadWithLink struct {
	Identifier SpreadIdentifier
}

type ProfitabilityPercent struct {
	Value     int64
	Precision int64
}
