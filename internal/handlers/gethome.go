package handlers

import (
    "go-sneed/internal/templates"
	"net/http"
)

type GetHomeHandler struct{}

func NewGetHomeHandler() *GetHomeHandler {
	return &GetHomeHandler{}
}

func (h *GetHomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c := templates.GuestIndex()
	//err := templates(c, "My website").Render(r.Context(), w)
    err := templates.Layout(c, "sneed").Render(r.Context(), w)

	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}
