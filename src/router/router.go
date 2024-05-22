package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/pbentes/80_20/src/middleware"
	"github.com/pbentes/80_20/src/views/components"
	"github.com/pbentes/80_20/src/views/fragments"
)

func SetupRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.Handle("GET /", middleware.ServeFiles(templ.Handler(components.Layout(fragments.Index))))

	router.HandleFunc("GET /index", templ.Handler(fragments.Index()).ServeHTTP)
	router.HandleFunc("GET /content", templ.Handler(fragments.Content()).ServeHTTP)
	return router
}
