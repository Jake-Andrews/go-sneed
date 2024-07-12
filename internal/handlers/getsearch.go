package handlers

import (
	store "go-sneed/internal/db"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"log"
	"net/http"
)

// Create a program to run with make that will add thumbnail/video paths
// to the database. possibly after migrate. tie in with tern?? for testing.
type GetSearchHandler struct {
    videoStore store.VideoStore
}

func NewGetSearchHandler(VideoStore store.VideoStore) *GetSearchHandler  {
    return &GetSearchHandler{videoStore: VideoStore}
}

func (h *GetSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    videos, err := h.videoStore.GetRandomVideos(r.Context(), 8)
    if err != nil {
        log.Printf("Error getting videos: %v", err)
        return
    }
    // hx-request, return partial
	//queryParams := r.URL.Query()
    //queryParams.Get("Hx-Request")
    if searchVal := r.Header.Get("Hx-Request"); searchVal != "" {
        if err := templates.Search(videos).Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Search(videos), r.Context(), w)
}

