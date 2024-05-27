package save_spreads

import (
	"encoding/json"
	"github.com/Ferum-Bot/HermesTrade/internal/platform/logger"
	dto "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/generated/schema"
	"github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/model"
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

func (handler *Handler) SaveSpreads(
	response http.ResponseWriter,
	r *http.Request,
) {
	request := dto.PutSpreadsStorageApiV1SaveSpreadsJSONRequestBody{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response.WriteHeader(http.StatusBadRequest)
		return
	}

	spreadsToSave := handler.convertSpreads(request.Spreads)

	_, err = handler.service.SaveSpreads(r.Context(), spreadsToSave)
	if err != nil {
		handler.logger.Errorf("Error searching spreads: %s", err.Error())

		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

func (handler *Handler) convertSpreads(
	spreads []dto.Spread,
) []model.Spread {
	result := make([]model.Spread, 0, len(spreads))

	for _, spread := range spreads {
		result = append(result, model.Spread{
			Identifier: model.SpreadIdentifier(spread.Identifier.String()),
		})
	}

	return result
}
