package service

import (
	dto "backend/dto/inscripcion"
	"backend/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

func CrearInscripcion(db *gorm.DB, req dto.CrearInscripcionDTO) error {
	// Verificar si ya está inscripto
	var existente model.Inscripcion
	if err := db.Where("usuario_id = ? AND actividad_id = ?", req.UsuarioID, req.ActividadID).First(&existente).Error; err == nil {
		return errors.New("el usuario ya está inscripto a esta actividad")
	}

	// Verificar si hay cupo
	var actividad model.Actividad
	if err := db.First(&actividad, req.ActividadID).Error; err != nil {
		return errors.New("actividad no encontrada")
	}

	var inscriptos int64
	db.Model(&model.Inscripcion{}).Where("actividad_id = ?", req.ActividadID).Count(&inscriptos)
	if int(inscriptos) >= actividad.Cupo {
		return errors.New("no hay cupo disponible para esta actividad")
	}

	// Crear inscripcion
	inscripcion := model.Inscripcion{
		UsuarioID:        req.UsuarioID,
		ActividadID:      req.ActividadID,
		FechaInscripcion: time.Now(),
	}

	if err := db.Create(&inscripcion).Error; err != nil {
		return err
	}

	return nil
}

// CancelarInscripcion cancela una inscripción existente
func CancelarInscripcion(db *gorm.DB, usuarioID, actividadID uint) error {
	var inscripcion model.Inscripcion
	if err := db.Where("usuario_id = ? AND actividad_id = ?", usuarioID, actividadID).First(&inscripcion).Error; err != nil {
		return errors.New("inscripción no encontrada")
	}

	if err := db.Delete(&inscripcion).Error; err != nil {
		return err
	}

	return nil
}

// ListarInscripciones lista las inscripciones de un usuario
func ListarInscripciones(db *gorm.DB, usuarioID uint) ([]dto.InscripcionDTO, error) {
	var inscripciones []model.Inscripcion
	if err := db.Where("usuario_id = ?", usuarioID).Find(&inscripciones).Error; err != nil {
		return nil, err
	}

	var result []dto.InscripcionDTO
	for _, insc := range inscripciones {
		var actividad model.Actividad
		if err := db.First(&actividad, insc.ActividadID).Error; err != nil {
			continue
		}

		result = append(result, dto.InscripcionDTO{
			UsuarioID:        insc.UsuarioID,
			ActividadID:      insc.ActividadID,
			FechaInscripcion: insc.FechaInscripcion,
			ActividadTitulo:  actividad.Titulo,
			Cupo:             actividad.Cupo,
		})
	}

	return result, nil
}
