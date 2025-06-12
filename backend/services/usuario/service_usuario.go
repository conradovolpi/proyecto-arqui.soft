package services

import (
	"os"

	"backend/clients/usuario"
	"backend/dto"
	"backend/models"
	"backend/utils"

	"golang.org/x/crypto/bcrypt"
)

type UsuarioService interface {
	Create(dto.UsuarioCreateDTO) (*dto.UsuarioResponseDTO, utils.ApiError)
	Login(dto.LoginDTO) (*dto.LoginResponseDTO, utils.ApiError)
	GetByID(uint) (*dto.UsuarioResponseDTO, utils.ApiError)
	GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError)
}

type usuarioService struct {
	client    usuario.UsuarioClientInterface
	jwtSecret string
}

func NewUsuarioService() UsuarioService {
	secret := os.Getenv("JWT_SECRET_KEY")
	if secret == "" {
		panic("JWT_SECRET_KEY no definida en entorno")
	}
	return &usuarioService{
		client:    usuario.UsuarioClient,
		jwtSecret: secret,
	}
}

func (s *usuarioService) Create(u dto.UsuarioCreateDTO) (*dto.UsuarioResponseDTO, utils.ApiError) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error al hashear la contraseña", err)
	}

	usuario := &models.Usuario{
		Nombre:   u.Nombre,
		Email:    u.Email,
		Password: string(hashed),
		Rol:      u.Rol,
	}

	if err := s.client.CreateUser(usuario); err != nil {
		return nil, utils.NewInternalServerApiError("Error creando usuario", err)
	}

	return &dto.UsuarioResponseDTO{
		UsuarioID: usuario.UsuarioID,
		Nombre:    usuario.Nombre,
		Email:     usuario.Email,
		Rol:       usuario.Rol,
	}, nil
}

func (s *usuarioService) Login(loginDTO dto.LoginDTO) (*dto.LoginResponseDTO, utils.ApiError) {
	usuario, err := s.client.GetByEmail(loginDTO.Email)
	if err != nil {
		return nil, utils.NewUnauthorizedApiError("Email o contraseña incorrectos")
	}

	if bcrypt.CompareHashAndPassword([]byte(usuario.Password), []byte(loginDTO.Password)) != nil {
		return nil, utils.NewUnauthorizedApiError("Email o contraseña incorrectos")
	}

	token := utils.GenerateJWT(usuario.UsuarioID, usuario.Rol, s.jwtSecret)

	return &dto.LoginResponseDTO{
		Token: token,
		Usuario: dto.UsuarioResponseDTO{
			UsuarioID: usuario.UsuarioID,
			Nombre:    usuario.Nombre,
			Email:     usuario.Email,
			Rol:       usuario.Rol,
		},
	}, nil
}

func (s *usuarioService) GetByID(id uint) (*dto.UsuarioResponseDTO, utils.ApiError) {
	usuario, err := s.client.GetByID(id)
	if err != nil {
		return nil, utils.NewNotFoundApiError("Usuario no encontrado")
	}

	return &dto.UsuarioResponseDTO{
		UsuarioID: usuario.UsuarioID,
		Nombre:    usuario.Nombre,
		Email:     usuario.Email,
		Rol:       usuario.Rol,
	}, nil
}

func (s *usuarioService) GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError) {
	usuarios, err := s.client.GetAll()
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error recuperando usuarios", err)
	}

	var dtos []dto.UsuarioResponseDTO
	for _, u := range usuarios {
		dtos = append(dtos, dto.UsuarioResponseDTO{
			UsuarioID: u.UsuarioID,
			Nombre:    u.Nombre,
			Email:     u.Email,
			Rol:       u.Rol,
		})
	}
	return dtos, nil
}
