package handlers

import (
	"encoding/json"
	"net/http"

	"RideUleam/internal/service/suscripciones"
)

type credenciales struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthHandler struct {
	auth *service.AuthService
}

func NewAuthHandler(auth *service.AuthService) *AuthHandler {
	return &AuthHandler{auth: auth}
}

func (h *AuthHandler) Registrar(w http.ResponseWriter, r *http.Request) {
	var creds credenciales

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	usuario, err := h.auth.Registrar(creds.Name, creds.Email, creds.Password)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusCreated, usuario)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds credenciales

	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		responderError(w, http.StatusBadRequest, "JSON inválido: "+err.Error())
		return
	}

	token, err := h.auth.Login(creds.Email, creds.Password)
	if err != nil {
		responderError(w, http.StatusBadRequest, err.Error())
		return
	}

	responderJSON(w, http.StatusOK, map[string]string{"token": token})
}
