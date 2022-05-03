package middlewares

import (
	"net/http"
	"strings"

	"github.com/waWanjohi/mux-auth/api/auth"
)

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth_header, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}

		tokenString := strings.TrimPrefix(auth_header, "Bearer ")

		claims, err := auth.GetClaimsFromToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r = r.WithContext(auth.SetJWTClaimsContext(r.Context(), claims))
		next.ServeHTTP(w, r)
	})
}

func SetMiddlewareAuthentication(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth_header := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth_header, "Bearer") {
			http.Error(w, "Not Authorized", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(auth_header, "Bearer ")
		claims, err := auth.GetClaimsFromToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		r = r.WithContext(auth.SetJWTClaimsContext(r.Context(), claims))
		next.ServeHTTP(w, r)
		next(w, r)
	}
}
