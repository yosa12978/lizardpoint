package middleware

import "net/http"

func Composition(
	f http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) http.Handler {
	for _, middleware := range middlewares {
		f = middleware(f)
	}
	return f
}
