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
		middleware.StripSlash(),
		middleware.Recovery(options.logger),
	)
	return handler
}

func addRoutes(router *http.ServeMux, opts options) {
	fileServer := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static", fileServer))

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		utils.RenderTemplate(w, "index", nil)
	})
}
