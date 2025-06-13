package services

import (
	"backend/clients/actividad"
	"backend/dto"
	"backend/models"
	"log"
)

type ActividadServiceInterface interface {
	CrearActividad(dto.ActividadCreateDTO) (dto.ActividadResponseDTO, error)
	GetByID(id uint) (dto.ActividadResponseDTO, error)
	GetAll() ([]dto.ActividadResponseDTO, error)
	Update(id uint, act dto.ActividadCreateDTO) error
	Delete(id uint) error
}

type actividadService struct {
	client actividad.ActividadClientInterface
}

func NewActividadService(c actividad.ActividadClientInterface) ActividadServiceInterface {
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

	actividad, err := s.client.Create(a)
	if err != nil {
		return dto.ActividadResponseDTO{}, err
	}

	return dto.ActividadResponseDTO{
		ActividadID:   actividad.ActividadID,
		HorarioInicio: actividad.HorarioInicio,
		HorarioFin:    actividad.HorarioFin,
		Titulo:        actividad.Titulo,
		Descripcion:   actividad.Descripcion,
		Instructor:    actividad.Instructor,
		Cupo:          actividad.Cupo,
		Categoria:     actividad.Categoria,
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
	log.Printf("Iniciando GetAll en el servicio de actividades")
	acts, err := s.client.GetAll()
	if err != nil {
		log.Printf("Error en servicio GetAll: %v", err)
		return nil, err
	}

	log.Printf("Actividades obtenidas del cliente: %d actividades", len(acts))
	var res []dto.ActividadResponseDTO
	for _, a := range acts {
		res = append(res, dto.ActividadResponseDTO{
			ActividadID:   a.ActividadID,
			HorarioInicio: a.HorarioInicio,
			HorarioFin:    a.HorarioFin,
			Titulo:        a.Titulo,
			Descripcion:   a.Descripcion,
			Instructor:    a.Instructor,
			Cupo:          a.Cupo,
			Categoria:     a.Categoria,
		})
	}
	log.Printf("Transformación de actividades completada")
	return res, nil
}

func (s *actividadService) Update(id uint, act dto.ActividadCreateDTO) error {
	a := models.Actividad{
		ActividadID:   id,
		HorarioInicio: act.HorarioInicio,
		HorarioFin:    act.HorarioFin,
		Titulo:        act.Titulo,
		Descripcion:   act.Descripcion,
		Instructor:    act.Instructor,
		Cupo:          act.Cupo,
		Categoria:     act.Categoria,
	}
	return s.client.Update(a)
}

func (s *actividadService) Delete(id uint) error {
	return s.client.Delete(id)
}
