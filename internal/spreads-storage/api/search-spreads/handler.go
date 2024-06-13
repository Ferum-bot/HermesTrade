package search_spreads

import (
	"encoding/json"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	dto "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type Handler struct {
	logger  logger.Logger
	service spreadService
}

func New(
	logger logger.Logger,
	service spreadService,
) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (handler *Handler) SearchSpreads(
	response http.ResponseWriter,
	r *http.Request,
	params dto.PostSpreadsStorageApiV1SearchSpreadsParams,
) {
	request := dto.PostSpreadsStorageApiV1SearchSpreadsJSONRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	offset := params.Offset
	limit := params.Limit
	filter := handler.parseFilter(request.ProfitabilityFilter, request.LengthFilter, request.FoundDateFilter)

	foundSpreads, err := handler.service.SearchSpreads(r.Context(), filter, offset, limit)
	if err != nil {
		handler.logger.Errorf("Error searching spreads: %s", err.Error())

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(response).Encode(handler.convertSpreadsToResponse(foundSpreads))
	if err != nil {
		handler.logger.Errorf("json.NewEncoder: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *Handler) parseFilter(
	profitabilityFilter *dto.SpreadProfitabilityFilter,
	lengthFilter *dto.SpreadLengthFilter,
	foundDateFilter *dto.SpreadFoundDateFilter,
) model.SpreadsFilter {
	filter := model.SpreadsFilter{}

	if profitabilityFilter != nil {
		filter.ProfitabilityFilter = &model.SpreadsProfitabilityFilter{}

		if profitabilityFilter.MaxProfitabilityPercent != nil {
			filter.ProfitabilityFilter.MaxProfitability = model.SpreadProfitabilityPercent{
				Value:     profitabilityFilter.MaxProfitabilityPercent.Value,
				Precision: profitabilityFilter.MaxProfitabilityPercent.Precision,
			}
		}

		if profitabilityFilter.MinProfitabilityPercent != nil {
			filter.ProfitabilityFilter.MinProfitability = model.SpreadProfitabilityPercent{
				Value:     profitabilityFilter.MinProfitabilityPercent.Value,
				Precision: profitabilityFilter.MinProfitabilityPercent.Precision,
			}
		}
	}

	if lengthFilter != nil {
		filter.LengthFilter = &model.SpreadsLengthFilter{}

		if lengthFilter.MinSpreadLength != nil {
			filter.LengthFilter.MinLength = *lengthFilter.MinSpreadLength
		}
		if lengthFilter.MaxSpreadLength != nil {
			filter.LengthFilter.MaxLength = *lengthFilter.MaxSpreadLength
		}
	}

	if foundDateFilter != nil {
		filter.FoundDateFilter = &model.SpreadsFoundDateFilter{}

		if foundDateFilter.StartFoundDate != nil {
			filter.FoundDateFilter.StartDate = *foundDateFilter.StartFoundDate
		}

		if foundDateFilter.EndFoundDate != nil {
			filter.FoundDateFilter.EndDate = *foundDateFilter.EndFoundDate
		}
	}

	return filter
}

func (handler *Handler) convertSpreadsToResponse(
	spreads []model.Spread,
) []dto.Spread {
	result := make([]dto.Spread, 0, len(spreads))

	for _, spread := range spreads {
		result = append(result, dto.Spread{
			Identifier: uuid.MustParse(string(spread.Identifier)),
			MetaInformation: struct {
				FoundAt              time.Time                      `json:"found_at"`
				ProfitabilityPercent dto.SpreadProfitabilityPercent `json:"profitability_percent"`
				SpreadLength         int64                          `json:"spread_length"`
			}(struct {
				FoundAt              time.Time
				ProfitabilityPercent dto.SpreadProfitabilityPercent
				SpreadLength         int64
			}{
				FoundAt: spread.MetaInformation.CreatedAt,
				ProfitabilityPercent: dto.SpreadProfitabilityPercent{
					Precision: spread.MetaInformation.ProfitabilityPercent.Precision,
					Value:     spread.MetaInformation.ProfitabilityPercent.Value,
				},
				SpreadLength: int64(spread.MetaInformation.Length),
			}),
			Elements: mapAssetCurrencyPairs(&spread.Head),
		})
	}

	return result
}

func mapAssetCurrencyPairs(startElement *model.SpreadElement) []dto.AssetCurrencyPair {
	result := make([]dto.AssetCurrencyPair, 0)

	currentElement := startElement

	for currentElement != nil {
		assetPair := currentElement.AssetPair

		identifier := uuid.New()

		result = append(result, dto.AssetCurrencyPair{
			Identifier: &identifier,
			BaseAsset: dto.Asset{
				ExternalIdentifier:  int64(assetPair.BaseAsset.ExternalIdentifier),
				UniversalIdentifier: int64(assetPair.BaseAsset.UniversalIdentifier),
			},
			QuotedAsset: dto.Asset{
				ExternalIdentifier:  int64(assetPair.QuotedAsset.ExternalIdentifier),
				UniversalIdentifier: int64(assetPair.QuotedAsset.UniversalIdentifier),
			},
			CurrencyRatio: dto.CurrencyRatio{
				Precision: assetPair.CurrencyRatio.Precision,
				Value:     assetPair.CurrencyRatio.Value,
			},
		})

		currentElement = currentElement.NextElement
	}

	return result
}
