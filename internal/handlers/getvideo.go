package handlers

import (
	store "go-sneed/internal/db"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type GetVideoHandler struct{
    db store.VideoStore
}

func NewGetVideoHandler(DB store.VideoStore) *GetVideoHandler {
    return &GetVideoHandler{db: DB}
}

func (h *GetVideoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    videoUUID := uuid.MustParse(r.FormValue("v"))
    log.Printf("Video UUID: %v\n", videoUUID)
    v, err := h.db.GetVideo(r.Context(), videoUUID)
    if err != nil {
        log.Printf("Error getting video from db: %v\nvideo_id: %v", err, videoUUID)
    }
    log.Printf("Video: %v\n", v)
    // hx-request, return partial
	queryParams := r.URL.Query()
    if searchVal := queryParams.Get("Hx-Request"); searchVal != "" {
        if err := templates.Video(v).Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Video(v), r.Context(), w)
}

