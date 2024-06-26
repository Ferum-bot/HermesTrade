package get_spreads

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

func (handler *Handler) GetSpreads(
	response http.ResponseWriter,
	r *http.Request,
) {
	request := dto.PostSpreadsStorageApiV1GetSpreadsJSONRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	spreadIDs := make([]model.SpreadIdentifier, 0, len(request.Identifiers))
	for _, identifier := range request.Identifiers {
		spreadIDs = append(spreadIDs, model.SpreadIdentifier(identifier))
	}

	spreads, err := handler.service.GetSpreadsWithLinks(r.Context(), spreadIDs)
	if err != nil {
		handler.logger.Errorf("Error getting spreads: %s", err.Error())

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(response).Encode(handler.convertSpreadsToResponse(spreads))
	if err != nil {
		handler.logger.Errorf("json.NewEncoder: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}

func (handler *Handler) convertSpreadsToResponse(
	spreads []model.SpreadWithLink,
) []dto.SpreadFull {
	result := make([]dto.SpreadFull, 0, len(spreads))

	for _, spread := range spreads {
		result = append(result, dto.SpreadFull{
			Identifier: uuid.New(),
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

func mapAssetCurrencyPairs(startElement *model.SpreadElementWithLink) []dto.AssetCurrencyPairWithLinks {
	result := make([]dto.AssetCurrencyPairWithLinks, 0)

	currentElement := startElement

	for currentElement != nil {
		assetPair := currentElement.AssetPair.AssetPair

		identifier := uuid.New()

		result = append(result, dto.AssetCurrencyPairWithLinks{
			Identifier: identifier,
			SourceLink: currentElement.AssetPair.SourceLink,
			PairLink:   currentElement.AssetPair.PairLink,
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
