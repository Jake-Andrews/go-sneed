package utils

import (
	"context"
	"net/http"

	"github.com/a-h/templ"
)

func RenderTemplate(templComponent templ.Component, ctx context.Context, w http.ResponseWriter) {
    if err := templComponent.Render(ctx, w); err != nil {
        http.Error(w, "error rendering template", http.StatusInternalServerError)
    }
}
