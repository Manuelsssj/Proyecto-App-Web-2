package service

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"

	models "RideUleam/internal/models/suscripciones"
)

type fakeUserRepoAuth struct {
	usuarios map[string]models.Usuario
}

func (f *fakeUserRepoAuth) CrearUsuario(u models.Usuario) (models.Usuario, error) {
	u.ID = len(f.usuarios) + 1
	f.usuarios[u.Email] = u
	return u, nil
}

func (f *fakeUserRepoAuth) BuscarUsuarioPorEmail(email string) (models.Usuario, bool) {
	u, ok := f.usuarios[email]
	return u, ok
}

func TestAuthService_RegistrarValido(t *testing.T) {
	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{}}
	srv := NewAuthService(repo)

	u, err := srv.Registrar("Usuario Test", "test@rideuleam.com", "123456")

	require.NoError(t, err)
	require.Equal(t, "Usuario Test", u.Name)
	require.Equal(t, "usuario", u.Rol)
	require.NotEmpty(t, u.PasswordHash)
}

func TestAuthService_RegistrarCredencialesInvalidas(t *testing.T) {
	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{}}
	srv := NewAuthService(repo)

	_, err := srv.Registrar("", "test@rideuleam.com", "123456")

	require.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_RegistrarEmailEnUso(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	require.NoError(t, err)

	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{
		"test@rideuleam.com": {
			ID:           1,
			Name:         "Usuario",
			Email:        "test@rideuleam.com",
			PasswordHash: string(hash),
			Rol:          "usuario",
		},
	}}
	srv := NewAuthService(repo)

	_, err = srv.Registrar("Otro", "test@rideuleam.com", "123456")

	require.ErrorIs(t, err, ErrEmailEnUso)
}

func TestAuthService_LoginValido(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	require.NoError(t, err)

	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{
		"test@rideuleam.com": {
			ID:           1,
			Name:         "Usuario",
			Email:        "test@rideuleam.com",
			PasswordHash: string(hash),
			Rol:          "usuario",
		},
	}}
	srv := NewAuthService(repo)

	token, err := srv.Login("test@rideuleam.com", "123456")

	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestAuthService_LoginPasswordIncorrecto(t *testing.T) {
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	require.NoError(t, err)

	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{
		"test@rideuleam.com": {
			ID:           1,
			Email:        "test@rideuleam.com",
			PasswordHash: string(hash),
			Rol:          "usuario",
		},
	}}
	srv := NewAuthService(repo)

	_, err = srv.Login("test@rideuleam.com", "incorrecta")

	require.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_LoginUsuarioNoExiste(t *testing.T) {
	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{}}
	srv := NewAuthService(repo)

	_, err := srv.Login("noexiste@rideuleam.com", "123456")

	require.ErrorIs(t, err, ErrCredencialesInvalidas)
}

func TestAuthService_GenerarYValidarToken(t *testing.T) {
	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{}}
	srv := NewAuthService(repo)

	token, err := srv.GenerarToken(models.Usuario{
		ID:    7,
		Email: "test@rideuleam.com",
		Rol:   "admin",
	})
	require.NoError(t, err)

	usuarioID, err := srv.ValidarToken(token)

	require.NoError(t, err)
	require.Equal(t, 7, usuarioID)
}

func TestAuthService_ValidarTokenInvalido(t *testing.T) {
	repo := &fakeUserRepoAuth{usuarios: map[string]models.Usuario{}}
	srv := NewAuthService(repo)

	_, err := srv.ValidarToken("token-invalido")

	require.ErrorIs(t, err, ErrCredencialesInvalidas)
}
