package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/middleware"
)

func newServer(ctx context.Context, addr string, logger logging.Logger) http.Server {
	router := http.NewServeMux()

	router.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, addr)
	})

	router.HandleFunc("GET /panic", func(w http.ResponseWriter, r *http.Request) {
		panic("unimplemented")
	})

	handler := middleware.Composition(
		router,
		middleware.Logger(logger),
		middleware.StripSlash,
		middleware.Recovery(logger),
	)

	return http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
