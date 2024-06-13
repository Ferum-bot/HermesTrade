package model

import "time"

type SpreadsFilter struct {
	ProfitabilityFilter *SpreadsProfitabilityFilter
	LengthFilter        *SpreadsLengthFilter
	FoundDateFilter     *SpreadsFoundDateFilter
}

type SpreadsProfitabilityFilter struct {
	MinProfitability SpreadProfitabilityPercent
	MaxProfitability SpreadProfitabilityPercent
}

type SpreadsLengthFilter struct {
	MinLength int64
	MaxLength int64
}

type SpreadsFoundDateFilter struct {
	StartDate time.Time
	EndDate   time.Time
}
