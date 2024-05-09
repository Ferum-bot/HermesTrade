package model

const DefaultMaxSpreadLength = int64(100)
const DefaultMinSpreadLength = int64(2)

type SpreadSearchSettings struct {
	MaxSpreadLength *int64
	MinSpreadLength *int64

	MinSearchProfitabilityRatio *SpreadProfitability
	MaxSearchProfitabilityRatio *SpreadProfitability
}
