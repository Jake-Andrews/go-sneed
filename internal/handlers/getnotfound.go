package handlers

import (
	"go-sneed/internal/templates"
	"net/http"
)

type GetNotFoundHandler struct{}

func NewGetNotFoundHandler() *GetNotFoundHandler {
	return &GetNotFoundHandler{}
}

func (h *GetNotFoundHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := templates.NotFound()
	err := templates.Layout(c, "Not Found").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}


