package services

import (
	"backend/clients/inscripcion"
	"backend/dto"
	"backend/models"
	"backend/utils"
	"time"
)

type InscripcionServiceInterface interface {
	Inscribir(dto.InscripcionCreateDTO) utils.ApiError
	Cancelar(dto.InscripcionCreateDTO) utils.ApiError
	GetPorUsuario(usuarioID uint) ([]dto.InscripcionResponseDTO, utils.ApiError)
	GetPorActividad(actividadID uint) ([]dto.InscripcionResponseDTO, utils.ApiError)
}

type inscripcionService struct {
	client inscripcion.InscripcionClientInterface
}

func NewInscripcionService(client inscripcion.InscripcionClientInterface) InscripcionServiceInterface {
	return &inscripcionService{client: client}
}

func (s *inscripcionService) Inscribir(d dto.InscripcionCreateDTO) utils.ApiError {
	existente, _ := s.client.Get(d.UsuarioID, d.ActividadID)
	if existente != nil {
		return utils.NewConflictApiError("Inscripción ya existente para este usuario y actividad")
	}

	insc := models.Inscripcion{
		UsuarioID:        d.UsuarioID,
		ActividadID:      d.ActividadID,
		FechaInscripcion: time.Now(),
	}

	if err := s.client.Create(insc); err != nil {
		return utils.NewInternalServerApiError("Error al crear inscripción", err)
	}

	return nil
}

func (s *inscripcionService) Cancelar(d dto.InscripcionCreateDTO) utils.ApiError {
	if err := s.client.Delete(d.UsuarioID, d.ActividadID); err != nil {
		return utils.NewInternalServerApiError("Error al cancelar inscripción", err)
	}
	return nil
}

func (s *inscripcionService) GetPorUsuario(usuarioID uint) ([]dto.InscripcionResponseDTO, utils.ApiError) {
	lista, err := s.client.GetByUsuario(usuarioID)
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error obteniendo inscripciones del usuario", err)
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

func (s *inscripcionService) GetPorActividad(actividadID uint) ([]dto.InscripcionResponseDTO, utils.ApiError) {
	lista, err := s.client.GetByActividad(actividadID)
	if err != nil {
		return nil, utils.NewInternalServerApiError("Error obteniendo inscripciones de la actividad", err)
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
