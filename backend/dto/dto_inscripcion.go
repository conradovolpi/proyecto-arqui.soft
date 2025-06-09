package dto

import "time"

type InscripcionCreateDTO struct {
	UsuarioID   uint `json:"usuario_id" binding:"required"`
	ActividadID uint `json:"actividad_id" binding:"required"`
}

type InscripcionResponseDTO struct {
	UsuarioID        uint      `json:"usuario_id"`
	ActividadID      uint      `json:"actividad_id"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
}
