package middleware

import (
	"context"
	"net/http"
	"strings"

	serviceMain "RideUleam/internal/service"
	service "RideUleam/internal/service/usuarioVehiculo"
)

// claveContexto es un tipo privado para la clave del context y evitar colisiones.
type claveContexto string

const ClaveRol claveContexto = "rol"

// ClaveUsuarioID es la clave bajo la que se guarda el ID del usuario autenticado.
const ClaveUsuarioID claveContexto = "usuarioID"

// Auth construye un middleware que exige un JWT valido en el header
// Authorization: Bearer <token>. Delega la validacion al AuthService: el
// middleware NO sabe de firmas ni de claims, solo de HTTP.
func Auth(auth *service.AuthService) func(http.Handler) http.Handler {
	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || !strings.EqualFold(partes[0], "Bearer") {
				responderNoAutorizado(w)
				return
			}

			claims, err := auth.ValidarClaims(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}

			ctx := context.WithValue(r.Context(), ClaveUsuarioID, claims.UsuarioID)
			ctx = context.WithValue(ctx, ClaveRol, claims.Rol)
			siguiente.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// AuthJWT adapta el servicio de autenticación incorporado por el módulo de
// rutas programadas al contexto que ya utiliza la rama de viajes inmediatos.
func AuthJWT(auth *serviceMain.AuthService) func(http.Handler) http.Handler {
	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			encabezado := r.Header.Get("Authorization")
			partes := strings.SplitN(encabezado, " ", 2)
			if len(partes) != 2 || !strings.EqualFold(partes[0], "Bearer") {
				responderNoAutorizado(w)
				return
			}

			claims, err := auth.ValidarToken(partes[1])
			if err != nil {
				responderNoAutorizado(w)
				return
			}

			ctx := context.WithValue(r.Context(), ClaveUsuarioID, claims.UsuarioID)
			ctx = context.WithValue(ctx, ClaveRol, claims.Rol)
			siguiente.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func responderNoAutorizado(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)
	_, _ = w.Write([]byte(`{"error":"token ausente o invalido"}`))
}

func RequiereRol(rolesPermitidos ...string) func(http.Handler) http.Handler {
	permitidos := map[string]bool{}
	for _, rol := range rolesPermitidos {
		permitidos[strings.ToLower(rol)] = true
	}

	return func(siguiente http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			rol, _ := r.Context().Value(ClaveRol).(string)
			if !permitidos[strings.ToLower(rol)] {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusForbidden)
				_, _ = w.Write([]byte(`{"error":"rol no autorizado"}`))
				return
			}
			siguiente.ServeHTTP(w, r)
		})
	}
}
