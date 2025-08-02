package middleware

import "net/http"

// Chain is a utility function to chain middlewares
func Chain(h http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, mw := range middlewares {
		h = mw(h)
	}
	return h
}
