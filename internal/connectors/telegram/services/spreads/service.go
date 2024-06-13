package spreads

import (
	"context"
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/internal/connectors/telegram/model"
	"github.com/Ferum-Bot/HermesTrade/internal/financial/constants"
	asset_pairs "github.com/Ferum-Bot/HermesTrade/internal/financial/providers/asset-pairs"
	currency_ratio "github.com/Ferum-Bot/HermesTrade/internal/financial/providers/currency-ratio"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type Service struct {
	assetPairs []model2.AssetCurrencyPair
}

func NewSpreadsService() *Service {
	return &Service{
		assetPairs: asset_pairs.ProvideAllAvailableCurrencyPairs(),
	}
}

func (service *Service) GetSpreads(
	ctx context.Context,
	profitability model.ProfitabilitySettingsType,
	parameters model.SpreadParameters,
) ([]model2.Spread, error) {
	switch profitability {
	case model.ProfitabilityAll:
		return service.generateSpreadsAll()
	case model.ProfitabilityPercent1:
		return service.generateSpreads1()
	case model.ProfitabilityPercent5:
		return service.generateSpreads5()
	case model.ProfitabilityPercent20:
		return service.generateSpreads20()
	}
	return service.generateSpreads5()
}

func (service *Service) generateSpreadsAll() ([]model2.Spread, error) {
	spreadsCount := rand.Int63()%10 + 4

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := service.generateSpread(model.ProfitabilityAll)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func (service *Service) generateSpreads1() ([]model2.Spread, error) {
	spreadsCount := rand.Int63()%8 + 4

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := service.generateSpread(model.ProfitabilityPercent1)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func (service *Service) generateSpreads5() ([]model2.Spread, error) {
	spreadsCount := rand.Int63() % 3

	spreads := make([]model2.Spread, 0, spreadsCount)
	for i := int64(0); i < spreadsCount; i++ {
		spread := service.generateSpread(model.ProfitabilityPercent5)
		spreads = append(spreads, spread)
	}

	return spreads, nil
}

func (service *Service) generateSpreads20() ([]model2.Spread, error) {
	return []model2.Spread{}, nil
}

func (service *Service) generateSpread(profitability model.ProfitabilitySettingsType) model2.Spread {
	spreadLength := rand.Int63()%4 + 3

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

	startIndex := rand.Int63() % int64(len(service.assetPairs))
	currentAssetPair := service.assetPairs[startIndex]

	for i := int64(0); i < spreadLength; i++ {
		newElement := &model2.SpreadElement{
			AssetPair: currentAssetPair,
		}

		if prevElement != nil {
			prevElement.NextElement = newElement
			prevElement = newElement
		} else {
			rootElement = newElement
			prevElement = newElement
		}

		if i == spreadLength-2 {
			firstAssetPair := service.assetPairs[startIndex]
			prevAssetPair := currentAssetPair

			firstName := constants.AssetNameByIdentifier[int64(prevAssetPair.QuotedAsset.UniversalIdentifier)]
			secondName := constants.AssetNameByIdentifier[int64(firstAssetPair.BaseAsset.UniversalIdentifier)]

			currentAssetPair = model2.AssetCurrencyPair{
				Identifier:    model2.AssetPairIdentifier(fmt.Sprintf("%s/%s", firstName, secondName)),
				BaseAsset:     prevAssetPair.QuotedAsset,
				QuotedAsset:   firstAssetPair.BaseAsset,
				CurrencyRatio: currency_ratio.ProvideCurrencyRatio(),
			}

		} else {
			nextAssetPairs := asset_pairs.ProvideAllNextAssetPair(service.assetPairs, currentAssetPair)
			nextIndex := rand.Int63() % int64(len(nextAssetPairs))

			currentAssetPair = nextAssetPairs[nextIndex]
		}
	}

	spread.Head = *rootElement

	switch profitability {
	case model.ProfitabilityAll:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 3,
			Value:     rand.Int63()%90 + 500,
		}
	case model.ProfitabilityPercent1:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 2,
			Value:     rand.Int63()%80 + 100,
		}
	case model.ProfitabilityPercent5:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 1,
			Value:     rand.Int63()%10 + 50,
		}
	default:
		spread.MetaInformation.ProfitabilityPercent = model2.SpreadProfitabilityPercent{
			Precision: 1,
			Value:     rand.Int63()%30 + 200,
		}
	}

	return spread
}
