package spreads

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Service struct {
}

func NewSpreadsService() *Service {
	return &Service{}
}

func (s *Service) GetSpreads(
	ctx context.Context,
	profitability model.ProfitabilitySettingsType,
	parameters model.SpreadParameters,
) ([]model2.Spread, error) {
	switch profitability {
	case model.ProfitabilityAll:
		return generateSpreadsAll()
	case model.ProfitabilityPercent1:
		return generateSpreads1()
	case model.ProfitabilityPercent5:
		return generateSpreads5()
	case model.ProfitabilityPercent20:
		return generateSpreads20()
	}
	return generateSpreads5()
}

func generateSpreadsAll() ([]model2.Spread, error) {
	spreadsCount := rand.Int63()%10 + 4

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := generateSpread(model.ProfitabilityAll)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func generateSpreads1() ([]model2.Spread, error) {
	spreadsCount := rand.Int63()%8 + 4

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := generateSpread(model.ProfitabilityPercent1)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func generateSpreads5() ([]model2.Spread, error) {
	spreadsCount := rand.Int63() % 3

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := generateSpread(model.ProfitabilityPercent5)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func generateSpreads20() ([]model2.Spread, error) {
	return []model2.Spread{}, nil
}

func generateSpread(profitability model.ProfitabilitySettingsType) model2.Spread {
	spreadLength := rand.Int63()%4 + 3

	availableAssetPairs := []model2.AssetCurrencyPair{
		{
			Identifier: "USD/EUR",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "EUR/USD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "USD/RUB",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "RUB/USD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "USD/AUD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "EUR/AUD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "GBP/AUD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    21123,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    1243,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "USD/TRY",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    3,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    13,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "TRY/USD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    98,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "TRY/USD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    283,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    32,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "USD/HUF",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    7,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    77,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "HUF/AUD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    21783,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "GBP/HUF",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 123,
				ExternalIdentifier:  0,
				SourceIdentifier:    234,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "HUF/USD",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 5,
				ExternalIdentifier:  0,
				SourceIdentifier:    3424,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "HUF/TRY",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 0,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
		{
			Identifier: "TRY/HUF",
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1678,
				ExternalIdentifier:  0,
				SourceIdentifier:    213,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 99,
				ExternalIdentifier:  0,
				SourceIdentifier:    123,
			},
			CurrencyRatio: generateCurrencyRation(),
		},
	}

	spread := model2.Spread{
		Identifier: model2.SpreadIdentifier(uuid.New().String()),
		MetaInformation: model2.SpreadMetaInformation{
			Length:    model2.SpreadLength(spreadLength),
			CreatedAt: time.Now(),
		},
	}

	var rootElement *model2.SpreadElement
	var prevElement *model2.SpreadElement
	prevElement = nil
	rootElement = nil

	for i := int64(0); i < spreadLength; i++ {
		index := rand.Int63() % int64(len(availableAssetPairs))

		newElement := &model2.SpreadElement{
			AssetPair: availableAssetPairs[index],
		}

		if prevElement != nil {
			prevElement.NextElement = newElement
			prevElement = newElement
		} else {
			rootElement = newElement
			prevElement = newElement
		}
	}

	spread.Head = *rootElement

	switch profitability {
	case model.ProfitabilityAll:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 3,
			Value:     rand.Int63()%30 + 5,
		}
	case model.ProfitabilityPercent1:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 2,
			Value:     rand.Int63()%80 + 1,
		}
	case model.ProfitabilityPercent5:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 1,
			Value:     rand.Int63()%50 + 50,
		}
	default:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 1,
			Value:     rand.Int63()%30 + 200,
		}
	}

	return spread
}

func generateCurrencyRation() model2.AssetsCurrencyRatio {
	precision := rand.Int63() % 3
	value := rand.Int63()%100000 + precision
	return model2.AssetsCurrencyRatio{
		Precision: precision,
		Value:     value,
	}
}
