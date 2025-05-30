package service

import (
    "gorm.io/gorm"
    "proyecto2025/backend/models"
)

// crear usuario
func CreateUsuario(db *gorm.DB, usuario *models.Usuario) error {
    return db.Create(usuario).Error
}

// editar usuario
func UpdateUsuario(db *gorm.DB, id uint, nuevosDatos *models.Usuario) error {
    var usuario models.Usuario
    if err := db.First(&usuario, id).Error; err != nil {
        return err
    }

    // Actualiza los campos necesarios
    usuario.Nombre = nuevosDatos.Nombre
    usuario.Email = nuevosDatos.Email
    usuario.Password = nuevosDatos.Password // Asegurate de que ya venga hasheado
    usuario.Rol = nuevosDatos.Rol

    return db.Save(&usuario).Error
}

//borrar usuairo
func DeleteUsuario(db *gorm.DB, id uint) error {
    return db.Delete(&models.Usuario{}, id).Error
}

//login


