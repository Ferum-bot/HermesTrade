package profitability_test

import (
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/profitability"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparer_ProfitabilityIsLessThan(t *testing.T) {
	type testCase struct {
		sourceProfitability model.SpreadProfitabilityPercent
		thanProfitability   model.SpreadProfitabilityPercent
		expectedResult      bool
	}

	tests := map[string]testCase{
		"similar_precision_easy": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6666,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     7777,
			},
			expectedResult: true,
		},
		"similar_precision_hard": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6666,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6671,
			},
			expectedResult: true,
		},
		"different_precision": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 2,
				Value:     112,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 4,
				Value:     6789,
			},
			expectedResult: false,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			comparer := profitability.NewComparer()

			actualResult := comparer.ProfitabilityIsLessThan(test.sourceProfitability, test.thanProfitability)

			assert.Equal(t, test.expectedResult, actualResult)
		})
	}
}

func TestComparer_ProfitabilityIsGreaterThan(t *testing.T) {
	type testCase struct {
		sourceProfitability model.SpreadProfitabilityPercent
		thanProfitability   model.SpreadProfitabilityPercent
		expectedResult      bool
	}

	tests := map[string]testCase{
		"similar_precision_easy": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6666,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     7777,
			},
			expectedResult: false,
		},
		"similar_precision_hard": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6666,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 3,
				Value:     6671,
			},
			expectedResult: false,
		},
		"different_precision": {
			sourceProfitability: model.SpreadProfitabilityPercent{
				Precision: 2,
				Value:     112,
			},
			thanProfitability: model.SpreadProfitabilityPercent{
				Precision: 4,
				Value:     6789,
			},
			expectedResult: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			comparer := profitability.NewComparer()

			actualResult := comparer.ProfitabilityIsGreaterThan(test.sourceProfitability, test.thanProfitability)

			assert.Equal(t, test.expectedResult, actualResult)
		})
	}
}
