package chat

import "time"

const (
	chatFieldChatID                    = "chat_id"
	chatFieldProfitabilitySettingsType = "profitability_settings_type"
	chatFieldUpdatedAt                 = "updated_at"
)

const (
	profitabilityTypeAll       = "profitability_all"
	profitabilityType1Percent  = "profitability_1_percent"
	profitabilityType5Percent  = "profitability_5_percent"
	profitabilityType20Percent = "profitability_20_percent"
)

type chatRow struct {
	ChatID                    int64     `bson:"chat_id"`
	ProfitabilitySettingsType string    `bson:"profitability_settings_type"`
	CreatedAt                 time.Time `bson:"created_at"`
	UpdatedAt                 time.Time `bson:"updated_at"`
}
