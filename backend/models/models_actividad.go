package models

import "time"

type Actividad struct {
	ActividadID   uint          `gorm:"primaryKey" json:"actividad_id"`
	HorarioInicio time.Time     `json:"horario_inicio"`
	HorarioFin    time.Time     `json:"horario_fin"`
	Titulo        string        `json:"titulo"`
	Descripcion   string        `json:"descripcion"`
	Instructor    string        `json:"instructor"`
	Cupo          int           `json:"cupo"`
	Categoria     string        `json:"categoria"`
	Inscripciones []Inscripcion `gorm:"foreignKey:ActividadID" json:"inscripciones,omitempty"`
}
