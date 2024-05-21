package message_converter

import (
	"fmt"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"strconv"
	"strings"
)

type Converter struct {
}

func NewMessageConverter() *Converter {
	return &Converter{}
}

func (c *Converter) ConvertSpreadsDataInOneMessage(
	spreads []model.Spread,
) string {
	resultString := strings.Builder{}

	resultString.WriteString(fmt.Sprintf("Found spreads count: %d\n", len(spreads)))

	for i := range spreads {
		spread := spreads[i]
		spreadTemplate := "%d. - [Profitability: %s][Length: %d][SpreadID: %s][Assets:%s]\n"

		resultString.WriteString(fmt.Sprintf(
			spreadTemplate, i+1,
			getSpreadProfitability(spread),
			getSpreadLength(spread),
			spread.Identifier,
			getSpreadsAssetsPair(spread),
		))
	}

	return resultString.String()
}

func getSpreadProfitability(spread model.Spread) string {
	profitabilityString := strings.Builder{}
	profitability := spread.MetaInformation.ProfitabilityPercent

	profitabilityValueString := strconv.FormatInt(profitability.Value, 64)

	if profitability.Precision >= int64(len(profitabilityValueString)) {
		profitabilityString.WriteString("0,")

		iterations := profitability.Precision - int64(len(profitabilityValueString))
		for i := 0; i < int(iterations); i++ {
			profitabilityString.WriteString("0")
		}
		profitabilityString.WriteString(profitabilityValueString)
	} else {
		iterations := -int64(len(profitabilityValueString)) - profitability.Precision
		for i := 0; i < int(iterations); i++ {
			profitabilityString.WriteString(string([]rune(profitabilityValueString)[i]))
		}

		profitabilityString.WriteString(",")

		for i := iterations; i < int64(len(profitabilityValueString)); i++ {
			profitabilityString.WriteString(string([]rune(profitabilityValueString)[i]))
		}
	}

	return profitabilityString.String()
}

func getSpreadLength(spread model.Spread) int64 {
	return int64(spread.MetaInformation.Length)
}

func getSpreadsAssetsPair(spread model.Spread) string {
	assetsString := strings.Builder{}
	length := spread.MetaInformation.Length

	currentElement := spread.Head
	for i := 0; i < int(length); i++ {
		assetPair := currentElement.AssetPair
		template := "Source(%d): AssetPair(%s): CurrencyRation(%s)"

		assetsString.WriteString(fmt.Sprintf(
			template,
			assetPair.BaseAsset.SourceIdentifier,
			assetPair.Identifier,
			getCurrencyRation(assetPair.CurrencyRatio),
		))

	}

	return assetsString.String()
}

func getCurrencyRation(ratio model.AssetsCurrencyRatio) string {
	ratioString := strings.Builder{}

	ratioValueString := strconv.FormatInt(ratio.Value, 64)

	if ratio.Precision >= int64(len(ratioValueString)) {
		ratioString.WriteString("0,")

		iterations := ratio.Precision - int64(len(ratioValueString))
		for i := 0; i < int(iterations); i++ {
			ratioString.WriteString("0")
		}
		ratioString.WriteString(ratioValueString)
	} else {
		iterations := -int64(len(ratioValueString)) - ratio.Precision
		for i := 0; i < int(iterations); i++ {
			ratioString.WriteString(string([]rune(ratioValueString)[i]))
		}

		ratioString.WriteString(",")

		for i := iterations; i < int64(len(ratioValueString)); i++ {
			ratioString.WriteString(string([]rune(ratioValueString)[i]))
		}
	}

	return ratioString.String()
}
