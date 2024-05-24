package search_spreads

import (
	dto "github.com/Ferum-Bot/HermesTrade/internal/spreads-storage/generated/schema"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (handler *Handler) SearchSpreads(
	response http.ResponseWriter,
	request *http.Request,
	params dto.PostSpreadsStorageApiV1SearchSpreadsParams,
) {

}
