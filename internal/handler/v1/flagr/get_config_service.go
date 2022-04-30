package flagr

import (
	"net/http"
	"strings"

	"go-microservice/internal/flagr"
	"go-microservice/internal/handler"
)

type Handler struct {
	flagr flagr.Service
}

func New(flagrService flagr.Service) *Handler {
	return &Handler{flagr: flagrService}
}

func (h *Handler) GetConfigsByService(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	service := strings.TrimSpace(r.URL.Query().Get("service"))
	if service == "" {
		res := handler.NewErrorResponse(handler.ErrMandatoryFieldsMissing, "tag and service name are mandatory fields", handler.ErrMandatoryFieldsMissing)
		res.Write(w, http.StatusBadRequest)
		return
	}

	configs, err := h.flagr.GetConfigByDescription(ctx, service)
	if err != nil {
		res := handler.NewErrorResponse(handler.ErrFailedToFetchFromFlagr, err.Error(), "flagr fetch failed")
		res.Write(w, http.StatusInternalServerError)
		return
	}

	res := handler.NewSuccessResponse(configs)
	res.Write(w, http.StatusOK)
}
