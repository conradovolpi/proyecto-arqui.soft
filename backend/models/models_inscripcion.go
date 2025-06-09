package models

import "time"

type Inscripcion struct {
	UsuarioID        uint      `gorm:"primaryKey" json:"usuario_id"`
	ActividadID      uint      `gorm:"primaryKey" json:"actividad_id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`

	Usuario   Usuario   `gorm:"foreignKey:UsuarioID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"usuario,omitempty"`
	Actividad Actividad `gorm:"foreignKey:ActividadID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"actividad,omitempty"`
}
