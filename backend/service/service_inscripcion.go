// service/service_inscripcion.go
package service

import (
	"errors"
	"time"

	"tu_modulo/dao"

	"gorm.io/gorm"
)

type InscripcionService struct {
	DB *gorm.DB
}

// CrearInscripcion crea una nueva inscripción, validando duplicados y cupo.
func (s *InscripcionService) CrearInscripcion(usuarioID, actividadID uint) error {
	// Verificar si ya está inscripto
	var existente dao.Inscripcion
	err := s.DB.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).First(&existente).Error
	if err == nil {
		return errors.New("el usuario ya está inscripto en esta actividad")
	}

	// Verificar existencia de actividad
	var actividad dao.Actividad
	if err := s.DB.First(&actividad, actividadID).Error; err != nil {
		return errors.New("actividad no encontrada")
	}

	// Verificar cupo
	var inscriptos int64
	s.DB.Model(&dao.Inscripcion{}).Where("actividad_id = ?", actividadID).Count(&inscriptos)
	if uint(inscriptos) >= actividad.Cupo {
		return errors.New("no hay cupos disponibles para esta actividad")
	}

	// Crear inscripción
	inscripcion := dao.Inscripcion{
		UsuarioID:        usuarioID,
		ActividadID:      actividadID,
		FechaInscripcion: time.Now(),
	}

	return s.DB.Create(&inscripcion).Error
}

// ObtenerTodas devuelve todas las inscripciones con relaciones
func (s *InscripcionService) ObtenerTodas() ([]dao.Inscripcion, error) {
	var inscripciones []dao.Inscripcion
	err := s.DB.Preload("Usuario").Preload("Actividad").Find(&inscripciones).Error
	return inscripciones, err
}

// ObtenerPorUsuario devuelve todas las inscripciones de un usuario
func (s *InscripcionService) ObtenerPorUsuario(usuarioID uint) ([]dao.Inscripcion, error) {
	var inscripciones []dao.Inscripcion
	err := s.DB.Preload("Actividad").Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error
	return inscripciones, err
}

// EliminarInscripcion elimina una inscripción existente
func (s *InscripcionService) EliminarInscripcion(usuarioID, actividadID uint) error {
	return s.DB.Delete(&dao.Inscripcion{}, "usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).Error
}
