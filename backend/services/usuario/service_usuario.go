package usuario

import (
	"backend/clients/usuario"
	"backend/dto"
	"backend/models"
	"backend/utils/errors"
	"crypto/md5"
	"encoding/hex"
)

type usuarioService struct{}

var UsuarioService usuarioServiceInterface

type usuarioServiceInterface interface {
	Create(dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, utils.ApiError)
	GetByID(uint) (dto.UsuarioResponseDTO, utils.ApiError)
	GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError)
}

func init() {
	UsuarioService = &usuarioService{}
}

func (s *usuarioService) Create(input dto.UsuarioCreateDTO) (dto.UsuarioResponseDTO, utils.ApiError) {
	hash := md5.Sum([]byte(input.Password))

	usuario := models.Usuario{
		Nombre:   input.Nombre,
		Email:    input.Email,
		Password: hex.EncodeToString(hash[:]),
		Rol:      input.Rol,
	}

	err := usuario.UsuarioClient.Create(&usuario)
	if err != nil {
		return dto.UsuarioResponseDTO{}, utils.NewInternalServerApiError("Error creando usuario", err)
	}

	return dto.UsuarioResponseDTO{
		UsuarioID: usuario.UsuarioID,
		Nombre:    usuario.Nombre,
		Email:     usuario.Email,
		Rol:       usuario.Rol,
	}, nil
}

func (s *usuarioService) GetByID(id uint) (dto.UsuarioResponseDTO, utils.ApiError) {
	u, err := usuario.UsuarioClient.GetByID(id)
	if err != nil {
		return dto.UsuarioResponseDTO{}, utils.NewNotFoundApiError("Usuario no encontrado")
	}

	return dto.UsuarioResponseDTO{
		UsuarioID: u.UsuarioID,
		Nombre:    u.Nombre,
		Email:     u.Email,
		Rol:       u.Rol,
	}, nil
}

func (s *usuarioService) GetAll() ([]dto.UsuarioResponseDTO, utils.ApiError) {
	us, err := usuario.UsuarioClient.GetAll()
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error listando usuarios", err)
	}

	var dtos []dto.UsuarioResponseDTO
	for _, u := range us {
		dtos = append(dtos, dto.UsuarioResponseDTO{
			UsuarioID: u.UsuarioID,
			Nombre:    u.Nombre,
			Email:     u.Email,
			Rol:       u.Rol,
		})
	}
	return dtos, nil
}
