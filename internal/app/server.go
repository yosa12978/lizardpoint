package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/yosa12978/lizardpoint/internal/logging"
)

func newServer(ctx context.Context, addr string, logger logging.Logger) http.Server {
	return http.Server{
		Addr: addr,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("incoming request")
			w.WriteHeader(200)
			fmt.Fprintf(w, addr)
		}),
	}
}
