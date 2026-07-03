package handlers

import (
	"encoding/json"
	"net/http"
)

func responderJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func responderError(w http.ResponseWriter, status int, mensaje string) {
	http.Error(w, mensaje, status)
}
