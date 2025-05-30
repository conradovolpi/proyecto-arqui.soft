package service

import (
    "gorm.io/gorm"
    "proyecto2025/backend/models"
)

// Obtener todas las actividades
func GetActividades(db *gorm.DB) ([]models.Actividad, error) {
    var actividades []models.Actividad
    if err := db.Find(&actividades).Error; err != nil {
        return nil, err
    }
    return actividades, nil
}

// Obtener una actividad por titulo
func GetActividadesPorTitulo(db *gorm.DB, titulo string) ([]models.Actividad, error) {
    var actividades []models.Actividad
    if err := db.Where("titulo = ?", titulo).Find(&actividades).Error; err != nil {
        return nil, err
    }
    return actividades, nil
}

// Crear nueva actividad
func CreateActividad(db *gorm.DB, actividad *models.Actividad) error {
    return db.Create(actividad).Error
}

// Actualizar actividad existente
func UpdateActividad(db *gorm.DB, id uint, nuevaActividad *models.Actividad) error {
    var actividad models.Actividad
    if err := db.First(&actividad, id).Error; err != nil {
        return err
    }
    return db.Model(&actividad).Updates(nuevaActividad).Error
}

// Eliminar actividad
func DeleteActividad(db *gorm.DB, id uint) error {
    return db.Delete(&models.Actividad{}, id).Error
}
