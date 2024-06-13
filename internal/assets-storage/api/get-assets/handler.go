package get_assets

import (
	"encoding/json"
	dto "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct {
	log           logger.Logger
	assetsService assetsService
}

func New(
	log logger.Logger,
	assetsService assetsService,
) *Handler {
	return &Handler{
		log:           log,
		assetsService: assetsService,
	}
}

func (handler *Handler) GetAssets(
	response http.ResponseWriter,
	r *http.Request,
	params dto.PostAssetsStorageApiV1GetAssetsParams,
) {
	request := dto.PostAssetsStorageApiV1GetAssetsJSONRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	offset := params.Offset
	limit := params.Limit
	assetFilters := handler.convertFilters(request.TimeFilter, request.SourceFilter, request.TypeFilter)

	foundAssets, err := handler.assetsService.GetAssets(r.Context(), assetFilters, offset, limit)
	if err != nil {
		handler.log.Errorf("GetAssets error: %s", err)
		response.WriteHeader(http.StatusInternalServerError)

		return
	}

	response.Header().Add("Content-Type", "application/json")

	err = json.NewEncoder(response).Encode(handler.convertAssetsToResponse(foundAssets))
	if err != nil {
		handler.log.Errorf("json.NewEncoder: %s", err)
		response.WriteHeader(http.StatusInternalServerError)
	}
}
func (handler *Handler) convertFilters(
	timeFilter *dto.TimeFilter,
	sourceFilter *dto.AssetSourceFilter,
	typeFilter *dto.AssetTypeFilter,
) model.AssetFilters {
	filter := model.AssetFilters{}

	if timeFilter != nil {
		filter.TimeFilter = &model.AssetTimeFilter{}
	}

	if sourceFilter != nil {
		filter.SourceFilter = &model.AssetSourceFilter{}
	}

	if typeFilter != nil {
		filter.TypeFilter = &model.AssetTypeFilter{}
	}

	return filter
}

func (handler *Handler) convertAssetsToResponse(assets []model.AssetCurrencyPair) []dto.AssetCurrencyPair {
	result := make([]dto.AssetCurrencyPair, 0, len(assets))

	for _, asset := range assets {
		result = append(result, dto.AssetCurrencyPair{
			Identifier: uuid.MustParse(string(asset.Identifier)),
			CurrencyRatio: dto.CurrencyRatio{
				Precision: asset.CurrencyRatio.Precision,
				Value:     asset.CurrencyRatio.Value,
			},
			BaseAsset: dto.Asset{
				ExternalIdentifier:  int64(asset.BaseAsset.ExternalIdentifier),
				UniversalIdentifier: int64(asset.BaseAsset.UniversalIdentifier),
			},
			QuotedAsset: dto.Asset{
				ExternalIdentifier:  int64(asset.QuotedAsset.ExternalIdentifier),
				UniversalIdentifier: int64(asset.QuotedAsset.UniversalIdentifier),
			},
		})
	}

	return result
}
