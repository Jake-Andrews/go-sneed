package handlers

import (
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
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
    utils.RenderTemplWithLayout(templates.Search(), r.Context(), w)
}
