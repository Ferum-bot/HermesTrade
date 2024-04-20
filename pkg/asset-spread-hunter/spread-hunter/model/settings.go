package model

import "time"

const DefaultMaxSpreadLength = int64(100)
const DefaultMinSpreadLength = int64(2)

const DefaultMaxSearchTimeDuration = 60 * time.Second
const DefaultMinSearchTimeDuration = 1 * time.Second

type SpreadSearchSettings struct {
	MaxSpreadLength *int64
	MinSpreadLength *int64

	MaxSearchTimeDuration *time.Duration
	MinSearchTimeDuration *time.Duration

	MinSearchProfitabilityRatio *SearchProfitabilityRatio
	MaxSearchProfitabilityRatio *SearchProfitabilityRatio
}

type SearchProfitabilityRatio struct {
	Precision int64
	Value     int64
}
