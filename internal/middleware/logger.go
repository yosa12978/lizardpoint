package middleware

import (
	"net/http"
	"time"

	"github.com/yosa12978/lizardpoint/internal/logging"
)

type wrappedWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *wrappedWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func Logger(logger logging.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			writer := &wrappedWriter{ResponseWriter: w, statusCode: 200}
			snapshot := time.Now()
			next.ServeHTTP(writer, r)
			latency := time.Since(snapshot).Microseconds()
			logger.Info("incoming request",
				"method", r.Method,
				"endpoint", r.URL.Path,
				"latency_us", latency,
				"status_code", writer.statusCode,
			)
		})
	}
}
