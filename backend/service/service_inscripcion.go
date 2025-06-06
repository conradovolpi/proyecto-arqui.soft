package service

import (
	"backend/dao"

	"gorm.io/gorm"
)

type InscripcionService struct {
	db *gorm.DB
}

func NewInscripcionService(db *gorm.DB) *InscripcionService {
	return &InscripcionService{db: db}
}

// CrearInscripcion intenta guardar una inscripción si no existe previamente
func (s *InscripcionService) CrearInscripcion(inscripcion *dao.Inscripcion) error {
	existe, err := dao.ExisteInscripcion(s.db, inscripcion.UsuarioID, inscripcion.ActividadID)
	if err != nil {
		return err
	}
	if existe {
		return nil // Ya existe, no se duplica
	}
	return dao.CrearInscripcion(s.db, inscripcion)
}

// ObtenerInscripcionesPorUsuario retorna todas las inscripciones de un usuario
func (s *InscripcionService) ObtenerInscripcionesPorUsuario(usuarioID uint) ([]dao.Inscripcion, error) {
	return dao.ObtenerInscripcionesPorUsuario(s.db, usuarioID)
}

// ObtenerInscripcionesPorActividad retorna todas las inscripciones de una actividad
func (s *InscripcionService) ObtenerInscripcionesPorActividad(actividadID uint) ([]dao.Inscripcion, error) {
	return dao.ObtenerInscripcionesPorActividad(s.db, actividadID)
}

// EliminarInscripcion elimina una inscripción
func (s *InscripcionService) EliminarInscripcion(usuarioID, actividadID uint) error {
	return dao.EliminarInscripcion(s.db, usuarioID, actividadID)
}
