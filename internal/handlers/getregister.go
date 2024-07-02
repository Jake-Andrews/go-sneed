package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetRegisterHandler struct{}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    utils.RenderTemplWithLayout(templates.RegisterPage(), r.Context(), w)
}
