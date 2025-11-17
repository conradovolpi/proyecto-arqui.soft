package services

import (
	"backend/clients"
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

// Helper function para calcular cupos disponibles
func (s *actividadService) calcularCuposDisponibles(actividadID uint, cupoTotal int) int {
	var inscripcionesCount int64
	clients.Db.Model(&models.Inscripcion{}).Where("actividad_id = ?", actividadID).Count(&inscripcionesCount)
	cuposDisponibles := cupoTotal - int(inscripcionesCount)
	if cuposDisponibles < 0 {
		cuposDisponibles = 0
	}
	return cuposDisponibles
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
		ActividadID:      actividad.ActividadID,
		HorarioInicio:    actividad.HorarioInicio,
		HorarioFin:       actividad.HorarioFin,
		Titulo:           actividad.Titulo,
		Descripcion:      actividad.Descripcion,
		Instructor:       actividad.Instructor,
		Cupo:             actividad.Cupo,
		CuposDisponibles: actividad.Cupo, // Nueva actividad, todos los cupos disponibles
		Categoria:        actividad.Categoria,
	}, nil
}

func (s *actividadService) GetByID(id uint) (dto.ActividadResponseDTO, error) {
	a, err := s.client.GetByID(id)
	if err != nil {
		return dto.ActividadResponseDTO{}, err
	}

	cuposDisponibles := s.calcularCuposDisponibles(a.ActividadID, a.Cupo)

	return dto.ActividadResponseDTO{
		ActividadID:      a.ActividadID,
		HorarioInicio:    a.HorarioInicio,
		HorarioFin:       a.HorarioFin,
		Titulo:           a.Titulo,
		Descripcion:      a.Descripcion,
		Instructor:       a.Instructor,
		Cupo:             a.Cupo,
		CuposDisponibles: cuposDisponibles,
		Categoria:        a.Categoria,
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
		cuposDisponibles := s.calcularCuposDisponibles(a.ActividadID, a.Cupo)
		res = append(res, dto.ActividadResponseDTO{
			ActividadID:      a.ActividadID,
			HorarioInicio:    a.HorarioInicio,
			HorarioFin:       a.HorarioFin,
			Titulo:           a.Titulo,
			Descripcion:      a.Descripcion,
			Instructor:       a.Instructor,
			Cupo:             a.Cupo,
			CuposDisponibles: cuposDisponibles,
			Categoria:        a.Categoria,
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
	// Verificar que la actividad existe antes de eliminar
	_, err := s.client.GetByID(id)
	if err != nil {
		return err
	}

	// Eliminar primero todas las inscripciones relacionadas
	if err := clients.Db.Where("actividad_id = ?", id).Delete(&models.Inscripcion{}).Error; err != nil {
		return err
	}

	// Luego eliminar la actividad
	return s.client.Delete(id)
}
