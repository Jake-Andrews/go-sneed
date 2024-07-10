package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type GetSearchHandler struct {}

func NewGetSearchHandler() *GetSearchHandler  {
    return &GetSearchHandler{}
}

func (h *GetSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    imagePaths, err := getTestPhotos()
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

func getTestPhotos() ([]string, error) {
	var photos []string

	// Read all files in the directory
	files, err := os.ReadDir("./static/images")
	if err != nil {
		return nil, err
	}

	// Iterate over each file and add its path to the photos slice if it's an image
	for _, file := range files {
		if !file.IsDir() {
			extension := filepath.Ext(file.Name())
			if extension == ".jpg" || extension == ".jpeg" || extension == ".png" || extension == ".gif" {
				photos = append(photos, filepath.Join("./static/images", file.Name()))
			}
		}
	}

	return photos, nil
}
