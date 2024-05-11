package cycles_spread_converter_test

import (
	"context"
	"github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/graph-algorithms/model"
	cycles_spread_converter "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/components/cycles-spread-converter"
	model2 "github.com/Ferum-Bot/HermesTrade/pkg/asset-spread-hunter/spread-hunter/model"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCyclesSpreadConverter_ConvertCyclesToSpreads(t *testing.T) {
	type testCase struct {
		cyclesIn        []model.GraphCycle
		allAssetsIn     []model2.AssetCurrencyPair
		expectedSpreads []model2.Spread
		expectedErr     error
	}

	assetsPairs := []model2.AssetCurrencyPair{
		// First Source
		{
			Identifier: model2.AssetPairIdentifier("EUR/USD"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  102,
				SourceIdentifier:    1,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  101,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1123,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier("USD/EUR"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  101,
				SourceIdentifier:    1,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  102,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 1,
				Value:     8,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 52,
				ExternalIdentifier:  202,
				SourceIdentifier:    1,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 51,
				ExternalIdentifier:  201,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier(uuid.New().String()),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 51,
				ExternalIdentifier:  201,
				SourceIdentifier:    1,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 52,
				ExternalIdentifier:  202,
				SourceIdentifier:    1,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1234,
			},
		},

		// Second Source
		{
			Identifier: model2.AssetPairIdentifier("USD/RUB"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  902,
				SourceIdentifier:    2,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  901,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 0,
				Value:     90,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier("RUB/USD"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  901,
				SourceIdentifier:    2,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 1,
				ExternalIdentifier:  902,
				SourceIdentifier:    2,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 3,
				Value:     1,
			},
		},

		// Third Source
		{
			Identifier: model2.AssetPairIdentifier("RUB/GBP"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  801,
				SourceIdentifier:    3,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  802,
				SourceIdentifier:    3,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 1,
				Value:     1,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier("GBP/RUB"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  802,
				SourceIdentifier:    3,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 3,
				ExternalIdentifier:  801,
				SourceIdentifier:    3,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 0,
				Value:     10,
			},
		},

		// Fours source
		{
			Identifier: model2.AssetPairIdentifier("GBP/EUR"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  702,
				SourceIdentifier:    4,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  701,
				SourceIdentifier:    4,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 1,
				Value:     1,
			},
		},
		{
			Identifier: model2.AssetPairIdentifier("EUR/GBP"),
			BaseAsset: model2.Asset{
				UniversalIdentifier: 2,
				ExternalIdentifier:  701,
				SourceIdentifier:    4,
			},
			QuotedAsset: model2.Asset{
				UniversalIdentifier: 4,
				ExternalIdentifier:  702,
				SourceIdentifier:    4,
			},
			CurrencyRatio: model2.AssetsCurrencyRatio{
				Precision: 0,
				Value:     1,
			},
		},
	}

	cycleVertexes := []model.GraphVertex{
		{
			Identifier: model.GraphVertexIdentifier(102),
		},
		{
			Identifier: model.GraphVertexIdentifier(101),
		},
		{
			Identifier: model.GraphVertexIdentifier(902),
		},
		{
			Identifier: model.GraphVertexIdentifier(901),
		},
		{
			Identifier: model.GraphVertexIdentifier(801),
		},
		{
			Identifier: model.GraphVertexIdentifier(802),
		},
		{
			Identifier: model.GraphVertexIdentifier(702),
		},
		{
			Identifier: model.GraphVertexIdentifier(701),
		},
	}

	spread := model2.Spread{
		Identifier: model2.SpreadIdentifier(uuid.New().String()),
		Head:       model2.SpreadElement{},
		MetaInformation: model2.SpreadMetaInformation{
			Length: model2.SpreadLength(4),
			ProfitabilityPercent: model2.SpreadProfitabilityPercent{
				Precision: 0,
				Value:     1,
			},
		},
	}

	tests := map[string]testCase{
		"empty_cycles": {
			cyclesIn:        []model.GraphCycle{},
			allAssetsIn:     assetsPairs,
			expectedSpreads: []model2.Spread{},
		},
		"one_simple_cycle": {
			cyclesIn: []model.GraphCycle{
				{
					Edges: []model.Edge{
						{
							SourceVertex: &cycleVertexes[0],
							TargetVertex: &cycleVertexes[1],
						},
						{
							SourceVertex: &cycleVertexes[1],
							TargetVertex: &cycleVertexes[2],
						},
						{
							SourceVertex: &cycleVertexes[2],
							TargetVertex: &cycleVertexes[3],
						},
						{
							SourceVertex: &cycleVertexes[3],
							TargetVertex: &cycleVertexes[4],
						},
						{
							SourceVertex: &cycleVertexes[4],
							TargetVertex: &cycleVertexes[5],
						},
						{
							SourceVertex: &cycleVertexes[5],
							TargetVertex: &cycleVertexes[6],
						},
						{
							SourceVertex: &cycleVertexes[6],
							TargetVertex: &cycleVertexes[7],
						},
						{
							SourceVertex: &cycleVertexes[7],
							TargetVertex: &cycleVertexes[0],
						},
					},
				},
			},
			allAssetsIn:     assetsPairs,
			expectedSpreads: []model2.Spread{spread},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			spreadConverter := cycles_spread_converter.NewCyclesSpreadConverter()

			actualSpreads, actualErr := spreadConverter.ConvertCyclesToSpreads(ctx, test.cyclesIn, test.allAssetsIn)
			if test.expectedErr != nil {
				assert.Error(t, actualErr)
				assert.Contains(t, actualErr.Error(), test.expectedErr.Error())
				return
			}

			assert.NoError(t, actualErr)
			assert.Equal(t, len(test.expectedSpreads), len(actualSpreads))
			for _, actualSpread := range actualSpreads {
				assertContainsSpread(t, test.expectedSpreads, actualSpread)
			}
		})
	}
}

func assertContainsSpread(t *testing.T, allSpreads []model2.Spread, targetSpread model2.Spread) {
	for _, spread := range allSpreads {
		if spreadsIsEqual(spread, targetSpread) {
			return
		}
	}

	t.Fatalf("all spreads(%v) not contains target spread(%v)", allSpreads, targetSpread)
}

func spreadsIsEqual(firstSpread model2.Spread, secondSpread model2.Spread) bool {
	if firstSpread.MetaInformation.Length != secondSpread.MetaInformation.Length {
		return false
	}
	if firstSpread.MetaInformation.ProfitabilityPercent != secondSpread.MetaInformation.ProfitabilityPercent {
		return false
	}

	return true
}
