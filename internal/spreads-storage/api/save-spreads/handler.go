package save_spreads

import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (handler *Handler) SaveSpreads(
	response http.ResponseWriter,
	request *http.Request,
) {

}
