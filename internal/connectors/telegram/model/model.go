package model

type ProfitabilitySettingsType int64

const (
	ProfitabilityAll ProfitabilitySettingsType = iota + 1
	ProfitabilityPercent1
	ProfitabilityPercent5
	ProfitabilityPercent20
)
