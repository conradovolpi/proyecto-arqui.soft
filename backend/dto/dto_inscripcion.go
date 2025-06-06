// dto/dto_inscripcion.go
package dto

type CrearInscripcionDTO struct {
	UsuarioID   uint `json:"usuario_id" binding:"required"`
	ActividadID uint `json:"actividad_id" binding:"required"`
}
