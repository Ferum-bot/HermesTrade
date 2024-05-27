package add_assets

import (
	"encoding/json"
	dto "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/assets-storage/model"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
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

func (handler *Handler) AddAssets(
	response http.ResponseWriter,
	r *http.Request,
	params dto.PutAssetsStorageApiV1AddAssetsParams,
) {
	request := dto.PutAssetsStorageApiV1AddAssetsJSONRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	sourceIdentifier := model.AssetSourceIdentifier(params.XHermesTradeAssetSourceIdentifier)
	assetsToAdd := handler.convertRequestAssets(request.AssetPairs, sourceIdentifier)

	_, err = handler.assetsService.AddAssets(r.Context(), assetsToAdd)
	if err != nil {
		handler.log.Errorf("Error adding assets: %s", err.Error())

		response.WriteHeader(http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (handler *Handler) convertRequestAssets(
	assets []dto.RequestAssetCurrencyPair,
	sourceIdentifier model.AssetSourceIdentifier,
) []model.AddAssetCurrencyPairData {
	result := make([]model.AddAssetCurrencyPairData, 0, len(assets))

	for _, asset := range assets {
		result = append(result, model.AddAssetCurrencyPairData{
			BaseAsset: model.Asset{
				SourceIdentifier:    sourceIdentifier,
				UniversalIdentifier: model.AssetUniversalIdentifier(asset.BaseAsset.UniversalIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(asset.BaseAsset.ExternalIdentifier),
			},
			QuotedAsset: model.Asset{
				SourceIdentifier:    sourceIdentifier,
				UniversalIdentifier: model.AssetUniversalIdentifier(asset.QuotedAsset.UniversalIdentifier),
				ExternalIdentifier:  model.AssetExternalIdentifier(asset.QuotedAsset.ExternalIdentifier),
			},
			CurrencyRatio: model.AssetCurrencyRatio{
				Value:     asset.CurrencyRatio.Value,
				Precision: asset.CurrencyRatio.Precision,
			},
		})
	}

	return result
}
