package middleware

import (
	"net/http"
	"strings"

	"RideUleam/internal/service"
)

func AuthJWT(authService *service.AuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				http.Error(w, "Token requerido", http.StatusUnauthorized)
				return
			}

			partes := strings.Split(header, " ")
			if len(partes) != 2 || partes[0] != "Bearer" {
				http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
				return
			}

			token := partes[1]

			_, err := authService.ValidarToken(token)
			if err != nil {
				http.Error(w, "Token inválido", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
