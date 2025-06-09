package services

import (
	"backend/clients/inscripcion"
	"backend/dto"
	"backend/models"
	"errors"
	"time"
)

type InscripcionServiceInterface interface {
	Inscribir(dto.InscripcionCreateDTO) error
	Cancelar(dto.InscripcionCreateDTO) error
	GetPorUsuario(usuarioID uint) ([]dto.InscripcionResponseDTO, error)
	GetPorActividad(actividadID uint) ([]dto.InscripcionResponseDTO, error)
}

type inscripcionService struct {
	client *inscripcion.InscripcionClient
}

func NewInscripcionService(c *inscripcion.InscripcionClient) InscripcionServiceInterface {
	return &inscripcionService{client: c}
}

func (s *inscripcionService) Inscribir(d dto.InscripcionCreateDTO) error {
	// validación sencilla: verificar si ya existe
	existente, _ := s.client.Get(d.UsuarioID, d.ActividadID)
	if existente != nil {
		return errors.New("usuario ya inscripto en esta actividad")
	}

	insc := models.Inscripcion{
		UsuarioID:        d.UsuarioID,
		ActividadID:      d.ActividadID,
		FechaInscripcion: time.Now(),
	}

	return s.client.Create(&insc)
}

func (s *inscripcionService) Cancelar(d dto.InscripcionCreateDTO) error {
	return s.client.Delete(d.UsuarioID, d.ActividadID)
}

func (s *inscripcionService) GetPorUsuario(usuarioID uint) ([]dto.InscripcionResponseDTO, error) {
	lista, err := s.client.GetByUsuario(usuarioID)
	if err != nil {
		return nil, err
	}

	var out []dto.InscripcionResponseDTO
	for _, i := range lista {
		out = append(out, dto.InscripcionResponseDTO{
			UsuarioID:        i.UsuarioID,
			ActividadID:      i.ActividadID,
			FechaInscripcion: i.FechaInscripcion,
		})
	}
	return out, nil
}

func (s *inscripcionService) GetPorActividad(actividadID uint) ([]dto.InscripcionResponseDTO, error) {
	lista, err := s.client.GetByActividad(actividadID)
	if err != nil {
		return nil, err
	}

	var out []dto.InscripcionResponseDTO
	for _, i := range lista {
		out = append(out, dto.InscripcionResponseDTO{
			UsuarioID:        i.UsuarioID,
			ActividadID:      i.ActividadID,
			FechaInscripcion: i.FechaInscripcion,
		})
	}
	return out, nil
}
