package service

import (
	dto "backend/dto/actividad"
	"backend/model"
	"errors"
	"time"

	"gorm.io/gorm"
)

// 1. Listar todas las actividades
func ListarActividades(db *gorm.DB) ([]dto.ActividadResponse, error) {
	var actividades []model.Actividad
	if err := db.Find(&actividades).Error; err != nil {
		return nil, err
	}

	var result []dto.ActividadResponse
	for _, a := range actividades {
		result = append(result, dto.ActividadResponse{
			ID:            a.ActividadID,
			Titulo:        a.Titulo,
			Descripcion:   a.Descripcion,
			HorarioInicio: a.HorarioInicio.Format(time.RFC3339),
			HorarioFin:    a.HorarioFin.Format(time.RFC3339),
			Instructor:    a.Instructor,
			Cupo:          uint(a.Cupo),
			Categoria:     a.Categoria,
		})
	}
	return result, nil
}

// 2. Obtener actividad por ID
func ObtenerActividadPorID(db *gorm.DB, id uint) (*dto.ActividadResponse, error) {
	var a model.Actividad
	if err := db.First(&a, id).Error; err != nil {
		return nil, errors.New("actividad no encontrada")
	}

	return &dto.ActividadResponse{
		ID:            a.ActividadID,
		Titulo:        a.Titulo,
		Descripcion:   a.Descripcion,
		HorarioInicio: a.HorarioInicio.Format(time.RFC3339),
		HorarioFin:    a.HorarioFin.Format(time.RFC3339),
		Instructor:    a.Instructor,
		Cupo:          uint(a.Cupo),
		Categoria:     a.Categoria,
	}, nil
}

// 3. Buscar actividades por palabra clave o categoría
func BuscarActividades(db *gorm.DB, query string) ([]dto.ActividadResponse, error) {
	var actividades []model.Actividad
	if err := db.Where("titulo LIKE ? OR categoria LIKE ?", "%"+query+"%", "%"+query+"%").Find(&actividades).Error; err != nil {
		return nil, err
	}

	var result []dto.ActividadResponse
	for _, a := range actividades {
		result = append(result, dto.ActividadResponse{
			ID:            a.ActividadID,
			Titulo:        a.Titulo,
			Descripcion:   a.Descripcion,
			HorarioInicio: a.HorarioInicio.Format(time.RFC3339),
			HorarioFin:    a.HorarioFin.Format(time.RFC3339),
			Instructor:    a.Instructor,
			Cupo:          uint(a.Cupo),
			Categoria:     a.Categoria,
		})
	}
	return result, nil
}

// cosas del final creacion update y baja de actividades

func CrearActividad(db *gorm.DB, req dto.ActividadRequest) error {
	inicio, err := time.Parse(time.RFC3339, req.HorarioInicio)
	if err != nil {
		return errors.New("formato de horario de inicio inválido")
	}
	fin, err := time.Parse(time.RFC3339, req.HorarioFin)
	if err != nil {
		return errors.New("formato de horario de fin inválido")
	}

	actividad := model.Actividad{
		Titulo:        req.Titulo,
		Descripcion:   req.Descripcion,
		HorarioInicio: inicio,
		HorarioFin:    fin,
		Instructor:    req.Instructor,
		Cupo:          int(req.Cupo),
		Categoria:     req.Categoria,
	}

	return db.Create(&actividad).Error
}

func ActualizarActividad(db *gorm.DB, id uint, req dto.ActividadRequest) error {
	var act model.Actividad
	if err := db.First(&act, id).Error; err != nil {
		return errors.New("actividad no encontrada")
	}

	inicio, err := time.Parse(time.RFC3339, req.HorarioInicio)
	if err != nil {
		return errors.New("formato de horario de inicio inválido")
	}
	fin, err := time.Parse(time.RFC3339, req.HorarioFin)
	if err != nil {
		return errors.New("formato de horario de fin inválido")
	}

	act.Titulo = req.Titulo
	act.Descripcion = req.Descripcion
	act.HorarioInicio = inicio
	act.HorarioFin = fin
	act.Instructor = req.Instructor
	act.Cupo = int(req.Cupo)
	act.Categoria = req.Categoria

	return db.Save(&act).Error
}

func EliminarActividad(db *gorm.DB, id uint) error {
	var act model.Actividad
	if err := db.First(&act, id).Error; err != nil {
		return errors.New("actividad no encontrada")
	}
	return db.Delete(&act).Error
}
