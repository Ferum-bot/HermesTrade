package chat

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
	ChatID                    int64  `bson:"chat_id"`
	ProfitabilitySettingsType string `bson:"profitability_settings_type"`
	CreatedAt                 int64  `bson:"created_at"`
	UpdatedAt                 int64  `bson:"updated_at"`
}
