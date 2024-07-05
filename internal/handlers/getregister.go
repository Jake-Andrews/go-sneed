package handlers

import (
	"go-sneed/internal/models"
	"go-sneed/internal/templates"
	"go-sneed/internal/utils"
	"net/http"
)

type GetRegisterHandler struct{}

func NewGetRegisterHandler() *GetRegisterHandler {
	return &GetRegisterHandler{}
}

func (h *GetRegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    formErrors := models.FormErrors{}
    formData := models.FormData{}
    utils.RenderTemplWithLayout(templates.RegisterPage(formErrors, formData), r.Context(), w)
}
