package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)
// Create a program to run with make that will add thumbnail/video paths
// to the database. possibly after migrate. tie in with tern?? for testing.
type GetSearchHandler struct {}

func NewGetSearchHandler() *GetSearchHandler  {
    return &GetSearchHandler{}
}

func (h *GetSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    imagePaths, err := getPhotos()
    if err != nil {
        log.Printf("Error getting image paths: %v", err)
        return
    }
    // hx-request, return partial
	//queryParams := r.URL.Query()
    //queryParams.Get("Hx-Request")
    if searchVal := r.Header.Get("Hx-Request"); searchVal != "" {
        if err := templates.Search(imagePaths).Render(r.Context(), w); err != nil {
            http.Error(w, "error", http.StatusInternalServerError)
        }
        return
    }
    //return full page
    utils.RenderTemplWithLayout(templates.Search(imagePaths), r.Context(), w)
}

func getPhotos() ([]string, error) {
	var photos []string

	err := filepath.Walk("./static/videos", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			extension := filepath.Ext(info.Name())
			if extension == ".jpg" || extension == ".jpeg" || extension == ".png" || extension == ".gif" {
				photos = append(photos, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return photos, nil
}

