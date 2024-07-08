package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetSearchHandler struct {}

func NewGetSearchHandler() *GetSearchHandler  {
    return &GetSearchHandler{}
}

func (h *GetSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.Search().Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Search(), r.Context(), w)
}
