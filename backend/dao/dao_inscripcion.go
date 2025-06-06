package dao

import "time"

type Inscripcion struct {
	UsuarioID        uint      `gorm:"primaryKey"`
	ActividadID      uint      `gorm:"primaryKey"`
	FechaInscripcion time.Time `gorm:"not null;autoCreateTime"`

	// Relaciones
	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;references:UsuarioID;constraint:OnDelete:CASCADE"`
	Actividad Actividad `gorm:"foreignKey:ActividadID;references:ActividadID;constraint:OnDelete:CASCADE"`
}
