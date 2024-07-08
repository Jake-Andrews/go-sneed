package handlers

import (
	"go-sneed/internal/templates"
	"log"
	"net/http"
)

type GetSearchHandler struct {}

func NewGetSearchHandler() *GetSearchHandler  {
    return &GetSearchHandler{}
}

func (h *GetSearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    log.Println(r)
    log.Println(r.URL)
    log.Println(r.Header)
    if err := templates.Search().Render(r.Context(), w); err != nil {
        http.Error(w, "sneed", http.StatusInternalServerError)
    }
}
