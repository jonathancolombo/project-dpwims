package auth

import (
	"net/http"
)

func RequireRole(requiredRole string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

			roleValue := request.Context().Value("role")
			if roleValue == nil {
				http.Error(writer, "missing role in token", http.StatusUnauthorized)
				return
			}

			role := roleValue.(string)

			if role != requiredRole {
				http.Error(writer, "forbidden: insufficient permissions", http.StatusForbidden)
				return
			}

			next.ServeHTTP(writer, request)
		})
	}
}
