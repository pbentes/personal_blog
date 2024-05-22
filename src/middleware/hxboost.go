package middleware

import (
	"net/http"

	"github.com/pbentes/80_20/src/templates"
)

func IsHXBoosted(template string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		template, _ := templates.GetTemplate(template, r.Header.Get("Hx-Boosted") == "")
		template.Render(r.Context(), w)
	})
}
