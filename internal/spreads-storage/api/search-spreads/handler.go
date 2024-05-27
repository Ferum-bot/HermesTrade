package search_spreads

import (
	"encoding/json"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	dto "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
	"github.com/google/uuid"
	"net/http"
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
			filter.ProfitabilityFilter.MaxProfitability = model.ProfitabilityPercent{
				Value:     profitabilityFilter.MaxProfitabilityPercent.Value,
				Precision: profitabilityFilter.MaxProfitabilityPercent.Precision,
			}
		}

		if profitabilityFilter.MinProfitabilityPercent != nil {
			filter.ProfitabilityFilter.MinProfitability = model.ProfitabilityPercent{
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
		})
	}

	return result
}
