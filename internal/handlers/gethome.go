package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetHomeHandler struct{}

func NewGetHomeHandler() *GetHomeHandler {
	return &GetHomeHandler{}
}

func (h *GetHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    utils.RenderTemplWithLayout(templates.GuestIndex(), r.Context(), w)
}
