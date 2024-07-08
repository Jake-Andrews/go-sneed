package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetTestHandler struct {}

func NewGetTestHandler() *GetTestHandler  {
    return &GetTestHandler{}
}

func (h *GetTestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.Test().Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Test(), r.Context(), w)
}
