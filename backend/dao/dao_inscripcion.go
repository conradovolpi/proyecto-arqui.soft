package dao

import (
	"time"

	"gorm.io/gorm"
)

type Inscripcion struct {
	UsuarioID        uint      `gorm:"primaryKey"`
	ActividadID      uint      `gorm:"primaryKey"`
	FechaInscripcion time.Time `gorm:"not null;autoCreateTime"`

	// Relaciones
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:UsuarioID;constraint:OnDelete:CASCADE"`
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ActividadID;constraint:OnDelete:CASCADE"`
}

// ===============================
// Funciones DAO
// ===============================

// CrearInscripcion guarda una nueva inscripción
func CrearInscripcion(db *gorm.DB, inscripcion *Inscripcion) error {
	return db.Create(inscripcion).Error
}

// ObtenerInscripcionesPorUsuario devuelve todas las inscripciones de un usuario
func ObtenerInscripcionesPorUsuario(db *gorm.DB, usuarioID uint) ([]Inscripcion, error) {
	var inscripciones []Inscripcion
	err := db.Where("usuario_id = ?", usuarioID).Preload("Actividad").Find(&inscripciones).Error
	return inscripciones, err
}

// ObtenerInscripcionesPorActividad devuelve todas las inscripciones a una actividad
func ObtenerInscripcionesPorActividad(db *gorm.DB, actividadID uint) ([]Inscripcion, error) {
	var inscripciones []Inscripcion
	err := db.Where("actividad_id = ?", actividadID).Preload("Usuario").Find(&inscripciones).Error
	return inscripciones, err
}

// EliminarInscripcion elimina una inscripción específica
func EliminarInscripcion(db *gorm.DB, usuarioID uint, actividadID uint) error {
	return db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).Delete(&Inscripcion{}).Error
}

// ExisteInscripcion verifica si un usuario ya está inscrito a una actividad
func ExisteInscripcion(db *gorm.DB, usuarioID uint, actividadID uint) (bool, error) {
	var count int64
	err := db.Model(&Inscripcion{}).
		Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).
		Count(&count).Error
	return count > 0, err
}
