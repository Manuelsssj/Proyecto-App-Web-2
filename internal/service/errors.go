package service

import "errors"

var ErrEmailEnUso = errors.New("email ya está en uso")
var ErrCredencialesInvalidas = errors.New("credenciales inválidas")
