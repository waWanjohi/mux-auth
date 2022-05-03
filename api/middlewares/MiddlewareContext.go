package middlewares

import (
	"net/http"
)

func SetMiddlewareContext(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// TODO: Modify context
		next(w, r)
	}
}
