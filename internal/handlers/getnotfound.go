package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetNotFoundHandler struct{}

func NewGetNotFoundHandler() *GetNotFoundHandler {
	return &GetNotFoundHandler{}
}

func (h *GetNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    utils.RenderTemplWithLayout(templates.NotFound(), r.Context(), w)
}


