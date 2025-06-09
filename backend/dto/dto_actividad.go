package dto

import "time"

type ActividadCreateDTO struct {
	HorarioInicio time.Time `json:"horario_inicio" binding:"required"`
	HorarioFin    time.Time `json:"horario_fin" binding:"required"`
	Titulo        string    `json:"titulo" binding:"required"`
	Descripcion   string    `json:"descripcion" binding:"required"`
	Instructor    string    `json:"instructor" binding:"required"`
	Cupo          int       `json:"cupo" binding:"required,gte=1"`
	Categoria     string    `json:"categoria" binding:"required"`
}

type ActividadResponseDTO struct {
	ActividadID   uint      `json:"actividad_id"`
	HorarioInicio time.Time `json:"horario_inicio"`
	HorarioFin    time.Time `json:"horario_fin"`
	Titulo        string    `json:"titulo"`
	Descripcion   string    `json:"descripcion"`
	Instructor    string    `json:"instructor"`
	Cupo          int       `json:"cupo"`
	Categoria     string    `json:"categoria"`
}
