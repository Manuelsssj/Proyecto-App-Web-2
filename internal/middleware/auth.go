package middleware

import (
	"context"
	"net/http"
	"strings"

	"RideUleam/internal/service"
)

type contextKey string

const claimsContextKey contextKey = "claims"

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

			claims, err := authService.ValidarToken(token)
			if err != nil {
				http.Error(w, "Token inválido", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), claimsContextKey, claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func ClaimsDesdeContext(ctx context.Context) (*service.Claims, bool) {
	claims, ok := ctx.Value(claimsContextKey).(*service.Claims)
	return claims, ok
}

func RolRequerido(rolesPermitidos ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			claims, ok := ClaimsDesdeContext(r.Context())
			if !ok {
				http.Error(w, "Token requerido", http.StatusUnauthorized)
				return
			}

			for _, rol := range rolesPermitidos {
				if claims.Rol == rol {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "No tienes permisos para acceder a este recurso", http.StatusForbidden)
		})
	}
}
