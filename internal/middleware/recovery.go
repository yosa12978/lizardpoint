package middleware

import (
	"fmt"
	"net/http"

	"github.com/yosa12978/lizardpoint/internal/logging"
)

func Recovery(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					logger.Error(
						"panic",
						"error", fmt.Sprintf("%+v", err),
						"method", r.Method,
						"endpoint", r.URL.Path,
					)
					w.Header().Set("Content-Type", "text/plain")
					w.WriteHeader(500)
					w.Write([]byte("Internal server error"))
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}
