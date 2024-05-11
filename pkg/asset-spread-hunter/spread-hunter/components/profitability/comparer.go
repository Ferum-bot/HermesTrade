package profitability

import (
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"strconv"
	"strings"
)

type Comparer struct {
}

func NewComparer() *Comparer {
	return &Comparer{}
}

func (c *Comparer) ProfitabilityIsLessThan(
	source model.SpreadProfitabilityPercent,
	than model.SpreadProfitabilityPercent,
) bool {
	biggestPrecision := source.Precision
	if than.Precision > biggestPrecision {
		biggestPrecision = than.Precision
	}

	sourceString := strings.Builder{}
	thanString := strings.Builder{}

	sourceString.WriteString(strconv.FormatInt(source.Value, 10))
	thanString.WriteString(strconv.FormatInt(than.Value, 10))

	for i := int64(0); i < biggestPrecision-source.Precision; i++ {
		sourceString.WriteString("0")
	}
	for i := int64(0); i < biggestPrecision-than.Precision; i++ {
		thanString.WriteString("0")
	}

	if sourceString.Len() != thanString.Len() {
		return sourceString.Len() < thanString.Len()
	}

	return sourceString.String() < thanString.String()
}

func (c *Comparer) ProfitabilityIsGreaterThan(
	source model.SpreadProfitabilityPercent,
	than model.SpreadProfitabilityPercent,
) bool {
	biggestPrecision := source.Precision
	if than.Precision > biggestPrecision {
		biggestPrecision = than.Precision
	}

	sourceString := strings.Builder{}
	thanString := strings.Builder{}

	sourceString.WriteString(strconv.FormatInt(source.Value, 10))
	thanString.WriteString(strconv.FormatInt(than.Value, 10))

	for i := int64(0); i < biggestPrecision-source.Precision; i++ {
		sourceString.WriteString("0")
	}
	for i := int64(0); i < biggestPrecision-than.Precision; i++ {
		thanString.WriteString("0")
	}

	if sourceString.Len() != thanString.Len() {
		return sourceString.Len() > thanString.Len()
	}

	return sourceString.String() > thanString.String()
}
