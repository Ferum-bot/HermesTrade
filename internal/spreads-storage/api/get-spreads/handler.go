package get_spreads

import "net/http"

type Handler struct {
}

func New() *Handler {
	return &Handler{}
}

func (handler *Handler) GetSpreads(
	response http.ResponseWriter,
	request *http.Request,
) {

}
