package middleware

import (
	"golang-playground/apitests/utils"
	"net/http"
)

// TokenMiddleware ...
func TokenMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("X-Token", utils.TokenGenerator())
		h.ServeHTTP(w, r)
	})
}
