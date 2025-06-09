package services

import (
	"backend/clients/actividad"
	"backend/dto"
	"backend/models"
	"errors"
)

type ActividadServiceInterface interface {
	CrearActividad(dto.ActividadCreateDTO) (dto.ActividadResponseDTO, error)
	GetByID(id uint) (dto.ActividadResponseDTO, error)
	GetAll() ([]dto.ActividadResponseDTO, error)
	Update(id uint, act dto.ActividadCreateDTO) error
	Delete(id uint) error
}

type actividadService struct {
	client *actividad.ActividadClient
}

func NewActividadService(c *actividad.ActividadClient) ActividadServiceInterface {
	return &actividadService{client: c}
}

func (s *actividadService) CrearActividad(d dto.ActividadCreateDTO) (dto.ActividadResponseDTO, error) {
	a := models.Actividad{
		HorarioInicio: d.HorarioInicio,
		HorarioFin:    d.HorarioFin,
		Titulo:        d.Titulo,
		Descripcion:   d.Descripcion,
		Instructor:    d.Instructor,
		Cupo:          d.Cupo,
		Categoria:     d.Categoria,
	}

	err := s.client.Create(&a)
	if err != nil {
		return dto.ActividadResponseDTO{}, err
	}

	return dto.ActividadResponseDTO{
		ActividadID:   a.ActividadID,
		HorarioInicio: a.HorarioInicio,
		HorarioFin:    a.HorarioFin,
		Titulo:        a.Titulo,
		Descripcion:   a.Descripcion,
		Instructor:    a.Instructor,
		Cupo:          a.Cupo,
		Categoria:     a.Categoria,
	}, nil
}

func (s *actividadService) GetByID(id uint) (dto.ActividadResponseDTO, error) {
	a, err := s.client.GetByID(id)
	if err != nil {
		return dto.ActividadResponseDTO{}, err
	}
	return dto.ActividadResponseDTO{
		ActividadID:   a.ActividadID,
		HorarioInicio: a.HorarioInicio,
		HorarioFin:    a.HorarioFin,
		Titulo:        a.Titulo,
		Descripcion:   a.Descripcion,
		Instructor:    a.Instructor,
		Cupo:          a.Cupo,
		Categoria:     a.Categoria,
	}, nil
}

func (s *actividadService) GetAll() ([]dto.ActividadResponseDTO, error) {
	acts, err := s.client.GetAll()
	if err != nil {
		return nil, err
	}
	var res []dto.ActividadRes
