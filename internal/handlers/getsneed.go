package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetSneedHandler struct {}

func NewGetSneedHandler() *GetSneedHandler  {
    return &GetSneedHandler{}
}

func (h *GetSneedHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    utils.RenderTemplWithLayout(templates.Sneed(), r.Context(), w)
}

