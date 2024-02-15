package middlewares

import (
	"fmt"
	"net/http"

	httpsetup "github.com/ankeshnirala/http-server-go1.22/cmd/api/rest/http-setup"
)

func LoggerMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("executing middleware before handler")

		next.ServeHTTP(w, r)

		fmt.Println("executing middleware after handler")
	})
}

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Reading token from headers
		tokenString := r.Header.Get("x-jwt-token")

		if tokenString == "" {
			httpsetup.WriteJSON(w, http.StatusUnauthorized, httpsetup.ApiError{Error: "unauthorized to access this resource"})
			return
		}

		next.ServeHTTP(w, r)
	})
}
