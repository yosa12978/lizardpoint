package app

import (
	"context"
	"net/http"

	"github.com/yosa12978/lizardpoint/internal/logging"
	"github.com/yosa12978/lizardpoint/internal/router"
)

func newServer(ctx context.Context, addr string, logger logging.Logger) http.Server {
	r := router.NewRouter(
		router.WithLogger(logger),
	)
	return http.Server{
		Addr:    addr,
		Handler: r,
	}
}
