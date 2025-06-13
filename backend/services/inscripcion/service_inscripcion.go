package services

import (
	"backend/clients"
	"backend/clients/inscripcion"
	"backend/dto"
	"backend/models"
	"backend/utils"
	"errors"
	"time"

	"gorm.io/gorm"
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
	// Verificar si el usuario existe
	var usuario models.Usuario
	if err := clients.Db.First(&usuario, d.UsuarioID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewNotFoundApiError("Usuario no encontrado")
		}
		return utils.NewInternalServerApiError("Error al verificar usuario", err)
	}

	// Verificar si la actividad existe
	var actividad models.Actividad
	if err := clients.Db.First(&actividad, d.ActividadID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewNotFoundApiError("Actividad no encontrada")
		}
		return utils.NewInternalServerApiError("Error al verificar actividad", err)
	}

	// Verificar si ya existe una inscripción
	existente, err := s.client.Get(d.UsuarioID, d.ActividadID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewInternalServerApiError("Error al verificar inscripción existente", err)
		}
	} else if existente != nil {
		return utils.NewConflictApiError("Ya estás inscrito en esta actividad")
	}

	// Verificar cupos disponibles (contando inscripciones existentes)
	inscripcionesExistentes, err := s.client.GetByActividad(d.ActividadID)
	if err != nil {
		return utils.NewInternalServerApiError("Error al obtener inscripciones existentes", err)
	}

	if len(inscripcionesExistentes) >= actividad.Cupo {
		return utils.NewConflictApiError("No hay cupos disponibles para esta actividad")
	}

	// Usar una transacción para asegurar la consistencia
	tx := clients.Db.Begin()
	if tx.Error != nil {
		return utils.NewInternalServerApiError("Error al iniciar la transacción", tx.Error)
	}

	// Crear la inscripción
	insc := models.Inscripcion{
		UsuarioID:        d.UsuarioID,
		ActividadID:      d.ActividadID,
		FechaInscripcion: time.Now(),
	}

	// Crear la inscripción dentro de la transacción
	if err := tx.Create(&insc).Error; err != nil {
		tx.Rollback()
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			return utils.NewConflictApiError("Ya estás inscrito en esta actividad")
		}
		return utils.NewInternalServerApiError("Error al crear la inscripción", err)
	}

	// Confirmar la transacción
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return utils.NewInternalServerApiError("Error al confirmar la inscripción", err)
	}

	return nil
}

func (s *inscripcionService) Cancelar(d dto.InscripcionCreateDTO) utils.ApiError {
	// Verificar si existe la inscripción
	existente, err := s.client.Get(d.UsuarioID, d.ActividadID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewNotFoundApiError("Inscripción no encontrada")
		}
		return utils.NewInternalServerApiError("Error al verificar inscripción", err)
	}

	if existente == nil {
		return utils.NewNotFoundApiError("Inscripción no encontrada")
	}

	// Obtener la actividad (necesario para la validación de existencia)
	var actividad models.Actividad
	if err := clients.Db.First(&actividad, d.ActividadID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return utils.NewNotFoundApiError("Actividad no encontrada")
		}
		return utils.NewInternalServerApiError("Error al verificar actividad", err)
	}

	// Usar una transacción para asegurar la consistencia
	tx := clients.Db.Begin()
	if tx.Error != nil {
		return utils.NewInternalServerApiError("Error al iniciar la transacción", tx.Error)
	}

	// Eliminar la inscripción
	if err := tx.Delete(&models.Inscripcion{}, "usuario_id = ? AND actividad_id = ?", d.UsuarioID, d.ActividadID).Error; err != nil {
		tx.Rollback()
		return utils.NewInternalServerApiError("Error al cancelar la inscripción", err)
	}

	// Confirmar la transacción
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return utils.NewInternalServerApiError("Error al confirmar la cancelación", err)
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
