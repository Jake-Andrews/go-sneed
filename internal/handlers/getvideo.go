package handlers

import (
	store "go-sneed/internal/db"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetVideoHandler struct{
    db store.VideoStore
}

func NewGetVideoHandler(DB store.VideoStore) *GetVideoHandler {
    return &GetVideoHandler{db: DB}
}

func (h *GetVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.Video().Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Video(), r.Context(), w)
}
