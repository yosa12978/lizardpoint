package middleware

import "net/http"

func Composition(f http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}
