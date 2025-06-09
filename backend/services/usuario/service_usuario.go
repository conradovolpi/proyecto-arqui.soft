package usuario

import (
	"backend/clients/usuario"
	"backend/dto"
	"backend/models"
	"crypto/md5"
	"encoding/hex"

	"github.com/golang-jwt/jwt/v5"
	log "github.com/sirupsen/logrus"
)

type usuarioService struct{}

var (
	UsuarioService UsuarioServiceInterface
)

type UsuarioServiceInterface interface {
	GetUsuarioByID(id uint) (dto.UsuarioResponseDTO, errores.ApiError)
	Login(loginDto dto.LoginDTO) (dto.TokenDto, errores.ApiError)
	CreateUsuario(dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, errores.ApiError)
}

func init() {
	UsuarioService = &usuarioService{}
}

var jwtKey = []byte("super_secret_key")

func (s *usuarioService) GetUsuarioByID(id uint) (dto.UsuarioResponseDTO, errores.ApiError) {
	user, err := usuario.UsuarioClient.GetByID(id)
	var userDto dto.UsuarioResponseDTO

	if err != nil || user == nil || user.UsuarioID == 0 {
		return userDto, errores.NewBadRequestApiError("usuario no encontrado")
	}

	userDto.UsuarioID = user.UsuarioID
	userDto.Nombre = user.Nombre
	userDto.Email = user.Email
	userDto.Rol = user.Rol
	return userDto, nil
}

func (s *usuarioService) Login(loginDto dto.LoginDTO) (dto.TokenDto, errores.ApiError) {
	log.Debug(loginDto)
	var tokenDto dto.TokenDto

	user, err := usuario.UsuarioClient.GetByEmail(loginDto.Email)
	if err != nil || user == nil {
		return tokenDto, errores.NewBadRequestApiError("Usuario no encontrado")
	}

	pswMd5 := md5.Sum([]byte(loginDto.Password))
	pswMd5string := hex.EncodeToString(pswMd5[:])

	if pswMd5string != user.Password {
		return tokenDto, errores.NewBadRequestApiError("Contraseña incorrecta")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UsuarioID,
	})

	tokenString, _ := token.SignedString(jwtKey)

	tokenDto.UsuarioID = user.UsuarioID
	tokenDto.Token = tokenString
	tokenDto.Rol = user.Rol

	return tokenDto, nil
}

func (s *usuarioService) CreateUsuario(registro dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, errores.ApiError) {
	var nuevo models.Usuario

	nuevo.Nombre = registro.Nombre
	nuevo.Email = registro.Email

	pswMd5 := md5.Sum([]byte(registro.Password))
	nuevo.Password = hex.EncodeToString(pswMd5[:])
	nuevo.Rol = registro.Rol

	nuevo, err := usuario.UsuarioClient.CreateUser(nuevo)
	if err != nil {
		return dto.UsuarioResponseDTO{}, errores.NewInternalServerApiError("Error al crear el usuario", err)
	}

	var registroResponse dto.UsuarioResponseDTO
	registroResponse.UsuarioID = nuevo.UsuarioID
	registroResponse.Nombre = nuevo.Nombre
	registroResponse.Email = nuevo.Email
	registroResponse.Rol = nuevo.Rol

	return registroResponse, nil
}
