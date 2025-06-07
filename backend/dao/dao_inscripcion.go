package dao

import (
	"time"

	"gorm.io/gorm"
)

type Inscripcion struct {
	gorm.Model
	UsuarioID        uint      `gorm:"not null"`
	ActividadID      uint      `gorm:"not null"`
	FechaInscripcion time.Time `gorm:"not null;autoCreateTime"`

	// Relaciones
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:ID"`
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ID"`
}

func (Inscripcion) TableName() string {
	return "inscripciones"
}
