package models

type Actividad struct {
	ActividadID   uint          `gorm:"primaryKey;autoIncrement"`
	HorarioInicio string        `gorm:"size:20;not null"`
	HorarioFin    string        `gorm:"size:20;not null"`
	Titulo        string        `gorm:"size:100;not null"`
	Descripcion   string        `gorm:"type:text"`
	Instructor    string        `gorm:"size:100;not null"`
	Duracion      int           `gorm:"not null"` // minutos
	Cupo          int           `gorm:"not null"`
	Categoria     string        `gorm:"size:50;not null"`
	FotoURL       string        `gorm:"size:255"` // Opcional, para la foto
	Inscripciones []Inscripcion `gorm:"foreignKey:ActividadID"`
}
