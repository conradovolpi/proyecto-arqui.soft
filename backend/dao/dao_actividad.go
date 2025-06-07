package dao

import (
	"time"

	"gorm.io/gorm"
)

type Actividad struct {
	gorm.Model
	HorarioInicio time.Time `gorm:"not null"`
	HorarioFin    time.Time `gorm:"not null"`
	Titulo        string    `gorm:"size:100;not null"`
	Descripcion   string    `gorm:"type:text"`
	Instructor    string    `gorm:"size:100;not null"`
	Duracion      int       `gorm:"not null"` // en minutos
	Cupo          int       `gorm:"not null"`
	Categoria     string    `gorm:"size:50;not null"`
	FotoURL       string    `gorm:"size:255"` // Opcional
}

func (Actividad) TableName() string {
	return "actividades"
}
