package middleware

import (
	"net/http"
)

func StripSlash() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			path := r.URL.Path
			if path == "" {
				r.URL.Path = "/"
			}
			if len(path) > 1 && path[len(path)-1] == '/' {
				r.URL.Path = path[:len(path)-1]
			}
			next.ServeHTTP(w, r)
		})
	}
}
