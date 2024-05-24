package get_assets

import (
	dto "github.com/Ferum-Bot/HermesTrade/internal/assets-storage/generated/schema"
	"net/http"
)

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (handler *Handler) GetAssets(
	response http.ResponseWriter,
	request *http.Request,
	params dto.PostAssetsStorageApiV1GetAssetsParams,
) {

}
