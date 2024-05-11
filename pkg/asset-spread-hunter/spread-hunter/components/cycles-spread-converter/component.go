package cycles_spread_converter

import (
	"context"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/platform/errors"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"strconv"
	"time"
)

const maxProfitabilityPrecision = 9

type CyclesSpreadConverter struct {
}

func NewCyclesSpreadConverter() *CyclesSpreadConverter {
	return &CyclesSpreadConverter{}
}

func (converter *CyclesSpreadConverter) ConvertCyclesToSpreads(
	ctx context.Context,
	cycles []model2.GraphCycle,
	sourceAssetPairs []model.AssetCurrencyPair,
) ([]model.Spread, error) {
	result := make([]model.Spread, 0, len(cycles))

	for _, cycle := range cycles {
		spread, err := converter.convertCycleToSpread(cycle, sourceAssetPairs)
		if err != nil {
			return nil, errors.Wrap(err, "converter.convertCycleToSpread")
		}

		result = append(result, spread)
	}

	return result, nil
}

func (converter *CyclesSpreadConverter) convertCycleToSpread(
	cycle model2.GraphCycle,
	sourceAssetPairs []model.AssetCurrencyPair,
) (model.Spread, error) {
	spread := model.Spread{
		Identifier: model.SpreadIdentifier(uuid.New().String()),
		MetaInformation: model.SpreadMetaInformation{
			Length:    model.SpreadLength(len(cycle.Edges) / 2),
			CreatedAt: time.Now(),
		},
	}

	spreadHead := converter.buildSpreadAssetsPairChain(cycle, sourceAssetPairs)
	spread.Head = spreadHead

	spreadProfitability, err := converter.calculateSpreadProfitability(spread)
	if err != nil {
		return model.Spread{}, errors.Wrap(err, "converter.calculateSpreadProfitability")
	}
	spread.MetaInformation.ProfitabilityPercent = spreadProfitability

	return spread, nil
}

func (converter *CyclesSpreadConverter) buildSpreadAssetsPairChain(
	cycle model2.GraphCycle,
	sourceAssetPairs []model.AssetCurrencyPair,
) model.SpreadElement {
	var rootElement *model.SpreadElement
	var currentElement *model.SpreadElement

	for _, edge := range cycle.Edges {
		assetPair := findAssetPairBy(*edge.SourceVertex, *edge.TargetVertex, sourceAssetPairs)
		if assetPair == nil {
			continue
		}

		if rootElement == nil {
			rootElement = &model.SpreadElement{
				AssetPair: *assetPair,
			}
			currentElement = rootElement
		} else {
			newElement := &model.SpreadElement{
				AssetPair: *assetPair,
			}
			currentElement.NextElement = newElement
			currentElement = newElement
		}
	}

	currentElement.NextElement = rootElement

	return *rootElement
}

func findAssetPairBy(
	sourceVertex model2.GraphVertex,
	targetVertex model2.GraphVertex,
	assetPairs []model.AssetCurrencyPair,
) *model.AssetCurrencyPair {
	var result *model.AssetCurrencyPair

	for _, assetPair := range assetPairs {
		baseAssetIdentifier := int64(assetPair.BaseAsset.ExternalIdentifier)
		quotedAssetIdentifiers := int64(assetPair.QuotedAsset.ExternalIdentifier)

		sourceVertexIdentifier := int64(sourceVertex.Identifier)
		targetVertexIdentifier := int64(targetVertex.Identifier)

		if baseAssetIdentifier == sourceVertexIdentifier &&
			quotedAssetIdentifiers == targetVertexIdentifier {
			result = &assetPair
			break
		}
	}

	return result
}

func (converter *CyclesSpreadConverter) calculateSpreadProfitability(
	spread model.Spread,
) (model.SpreadProfitabilityPercent, error) {
	rootElement := spread.Head
	allSpreadAssetPairs := make([]model.AssetCurrencyPair, 0, spread.MetaInformation.Length)

	currentElement := rootElement
	for {
		allSpreadAssetPairs = append(allSpreadAssetPairs, currentElement.AssetPair)

		if currentElement.NextElement != nil {
			currentElement = *currentElement.NextElement
		}

		if currentElement.AssetPair.Identifier == rootElement.AssetPair.Identifier {
			break
		}
	}

	commonPrecision := findCommonPrecision(allSpreadAssetPairs)
	initialTradeCurrencyValue := initTradeCurrencyValueFromPrecision(commonPrecision)
	currentTradeCurrencyValue := initialTradeCurrencyValue
	currentTradeCurrencyPrecision := int64(0)

	var err error
	for currentIndex := range allSpreadAssetPairs {
		assetPair := allSpreadAssetPairs[currentIndex]
		currentTradeCurrencyValue, currentTradeCurrencyPrecision, err = calculateNextCurrencyValueFrom(
			currentTradeCurrencyValue, currentTradeCurrencyPrecision, assetPair.CurrencyRatio,
		)
		if err != nil {
			return model.SpreadProfitabilityPercent{}, errors.Wrap(err, "calculateNextCurrencyValueFrom")
		}
	}

	profitability := calculateProfitabilityPercent(
		initialTradeCurrencyValue,
		currentTradeCurrencyValue,
		currentTradeCurrencyPrecision,
	)
	return profitability, nil
}

func findCommonPrecision(assetPairs []model.AssetCurrencyPair) int64 {
	commonPrecision := int64(0)
	for _, pair := range assetPairs {
		if pair.CurrencyRatio.Precision > commonPrecision {
			commonPrecision = pair.CurrencyRatio.Precision
		}
	}

	if commonPrecision > maxProfitabilityPrecision {
		commonPrecision = maxProfitabilityPrecision
	}

	return commonPrecision
}

func initTradeCurrencyValueFromPrecision(precision int64) int64 {
	value := int64(1)

	for i := int64(0); i < precision; i++ {
		value *= 10
	}

	return value
}

func calculateNextCurrencyValueFrom(
	currentCurrencyValue int64,
	currentCurrencyPrecision int64,
	currencyRation model.AssetsCurrencyRatio,
) (int64, int64, error) {
	resultValue := currentCurrencyValue * currencyRation.Value
	resultValueString := strconv.FormatInt(resultValue, 10)

	commonPrecision := currencyRation.Precision + currentCurrencyPrecision
	if commonPrecision > maxProfitabilityPrecision {
		commonPrecision = maxProfitabilityPrecision
	}

	for len(resultValueString) > 0 && commonPrecision > 0 {
		if resultValueString[len(resultValueString)-1] == '0' {
			commonPrecision--
			resultValueString = resultValueString[:len(resultValueString)-1]
		} else {
			break
		}
	}
	if len(resultValueString) == 0 {
		commonPrecision = 0
	}

	resultValue, err := strconv.ParseInt(resultValueString, 10, 64)
	if err != nil {
		return 0, 0, errors.Wrap(err, "strconv.ParseInt")
	}

	return resultValue, commonPrecision, nil
}

func calculateProfitabilityPercent(
	initialTradeCurrencyValue int64,
	resultTradeCurrencyValue int64,
	resultTradeCurrencyPrecision int64,
) model.SpreadProfitabilityPercent {
	for i := int64(0); i < resultTradeCurrencyPrecision; i++ {
		initialTradeCurrencyValue *= 10
	}

	profitabilityValue := resultTradeCurrencyValue - initialTradeCurrencyValue
	if profitabilityValue <= 0 {
		return model.SpreadProfitabilityPercent{
			Precision: 0,
			Value:     0,
		}
	}

	profitabilityPercent := float64(profitabilityValue) / float64(initialTradeCurrencyValue) * 100

	return model.SpreadProfitabilityPercent{
		Precision: 0,
		Value:     int64(profitabilityPercent),
	}
}
