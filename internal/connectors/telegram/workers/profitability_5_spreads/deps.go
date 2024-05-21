package profitability_5_spreads

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
)

type spreadsService interface {
	GetSpreads(
		ctx context.Context,
		profitability model2.ProfitabilitySettingsType,
		parameters model2.SpreadParameters,
	) ([]model.Spread, error)
}

type spreadMessageConverter interface {
	ConvertSpreadsDataInOneMessage(
		[]model.Spread,
	) string
}

type chatService interface {
	GetChatsCountWithProfitability(
		ctx context.Context,
		profitability model2.ProfitabilitySettingsType,
	) (int64, error)

	GetChatsWithProfitability(
		ctx context.Context,
		profitability model2.ProfitabilitySettingsType,
		pageNumber, pageSize int64,
	) ([]model2.Chat, error)
}

type telegramBot interface {
	SendMessage(
		ctx context.Context,
		chatID model2.ChatID,
		message string,
	) error
}
