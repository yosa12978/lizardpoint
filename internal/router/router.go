package router

import (
	"net/http"

	"github.com/yosa12978/lizardpoint/internal/middleware"
	"github.com/yosa12978/lizardpoint/pkg/utils"
)

func NewRouter(opts ...optionFunc) http.Handler {
	options := newOptions(opts...)
	router := http.NewServeMux()
	addRoutes(router, options)
	handler := middleware.Composition(
		router,
		middleware.Logger(options.logger),
		middleware.StripSlash,
		middleware.Recovery(options.logger),
	)
	return handler
}

func addRoutes(router *http.ServeMux, opts options) {
	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		utils.RenderTemplate(w, "index", nil)
	})

	router.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("unimplemented")
	})
}
