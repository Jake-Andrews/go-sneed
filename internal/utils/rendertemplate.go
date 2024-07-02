package utils

import (
	"context"
	"go-sneed/internal/templates"
	"net/http"

	"github.com/a-h/templ"
)

func RenderTemplWithLayout(c templ.Component, ctx context.Context, w http.ResponseWriter) {
    if err := templates.Layout(c, "sneed").Render(ctx, w); err != nil {
        http.Error(w, "error rendering template", http.StatusInternalServerError)
    }
}
