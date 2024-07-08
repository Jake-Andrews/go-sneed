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
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.NotFound().Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.NotFound(), r.Context(), w)
}


