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
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.GuestIndex().Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.GuestIndex(), r.Context(), w)
}
