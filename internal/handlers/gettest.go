package handlers

import (
	"go-sneed/internal/templates"
	"net/http"
)

type GetTestHandler struct {}

func NewGetTestHandler() *GetTestHandler  {
    return &GetTestHandler{}
}

func (h *GetTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    if err := templates.Test().Render(r.Context(), w); err != nil {
        http.Error(w, "sneed", http.StatusInternalServerError)
    }
}
