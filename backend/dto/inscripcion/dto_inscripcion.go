package dto

import "time"

//nose si lo usamos a este
type InscripcionDTO struct {
	UsuarioID        uint      `json:"usuario_id"`
	ActividadID      uint      `json:"actividad_id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
	ActividadTitulo  string    `json:"actividad_titulo"`
	Cupo             int       `json:"cupo"`
}

type CrearInscripcionDTO struct {
	UsuarioID   uint `json:"usuario_id" binding:"required"`
	ActividadID uint `json:"actividad_id" binding:"required"`
}
