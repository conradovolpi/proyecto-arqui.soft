package services

import (
	"backend/clients/usuario"
	"backend/dto"
	"backend/models"
	"crypto/md5"
	"encoding/hex"
	"errors"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("super_secret_key")

type UsuarioServiceInterface interface {
	CrearUsuario(dto dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, error)
	Login(email, password string) (dto.TokenDto, error)
	GetByID(id uint) (dto.UsuarioResponseDTO, error)
}

type usuarioService struct {
	client *usuario.UsuarioClient
}

func NewUsuarioService(c *usuario.UsuarioClient) UsuarioServiceInterface {
	return &usuarioService{client: c}
}

func (s *usuarioService) CrearUsuario(dto dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, error) {
	hashed := md5.Sum([]byte(dto.Password))
	password := hex.EncodeToString(hashed[:])

	user := models.Usuario{
		Nombre:   dto.Nombre,
		Email:    dto.Email,
		Password: password,
		Rol:      dto.Rol,
	}

	err := s.client.Create(&user)
	if err != nil {
		return dto.UsuarioResponseDTO{}, err
	}

	return dto.UsuarioResponseDTO{
		UsuarioID: user.UsuarioID,
		Nombre:    user.Nombre,
		Email:     user.Email,
		Rol:       user.Rol,
	}, nil
}

func (s *usuarioService) Login(email, password string) (dto.TokenDto, error) {
	user, err := s.client.GetByEmail(email)
	if err != nil || user == nil {
		return dto.TokenDto{}, errors.New("usuario no encontrado")
	}

	passwordHash := md5.Sum([]byte(password))
	passwordHex := hex.EncodeToString(passwordHash[:])

	if user.Password != passwordHex {
		return dto.TokenDto{}, errors.New("contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UsuarioID,
	})

	tokenString, _ := token.SignedString(jwtKey)

	return dto.TokenDto{
		Token:     tokenString,
		UsuarioID: user.UsuarioID,
		Rol:       user.Rol,
	}, nil
}

func (s *usuarioService) GetByID(id uint) (dto.UsuarioResponseDTO, error) {
	user, err := s.client.GetByID(id)
	if err != nil || user == nil {
		return dto.UsuarioResponseDTO{}, errors.New("usuario no encontrado")
	}

	return dto.UsuarioResponseDTO{
		UsuarioID: user.UsuarioID,
		Nombre:    user.Nombre,
		Email:     user.Email,
		Rol:       user.Rol,
	}, nil
}
