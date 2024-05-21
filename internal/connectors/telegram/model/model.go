package model

import "time"

type ProfitabilitySettingsType int64

const (
	ProfitabilityAll ProfitabilitySettingsType = iota + 1
	ProfitabilityPercent1
	ProfitabilityPercent5
	ProfitabilityPercent20
)

type Chat struct {
	ChatID            ChatID
	ProfitabilityType ProfitabilitySettingsType
}

type SpreadParameters struct {
	MaxLength int64
	FoundTime time.Time
}
